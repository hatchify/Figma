package figma

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Hatch1fy/errors"
)

const (
	baseURL    = "https://api.figma.com"
	filePath   = "v1/files/%s"
	nodesPath  = "v1/files/%s/nodes"
	stylesPath = "v1/teams/%s/styles"
)

const (
	// ErrEmptyAccessToken will return when provided access token is empty
	ErrEmptyAccessToken = errors.Error("provided access token is empty")

	// ErrFillIsGradient will return when a gradient is found and should not be stored
	ErrFillIsGradient = errors.Error("fill isgradient")
	// ErrFillIsSolid will return when a solid is found and should not be stored
	ErrFillIsSolid = errors.Error("fill is solid")
)

// New will return a new instance of the HubSpot manager
func New(accessToken, teamID string) (fp *Figma, err error) {
	var f Figma
	// f.out = journaler.New("Figma")
	if accessToken == "" {
		err = ErrEmptyAccessToken
		return
	}

	f.AccessToken = accessToken
	fp = &f
	return
}

// NewFromFile will return a new auth pair from a provided filename
func NewFromFile(fpath string) (fp *Figma, err error) {
	var f Figma
	// f.out = journaler.New("Figma")
	ext := filepath.Ext(fpath)
	switch ext {
	case ".toml":
		if _, err = toml.DecodeFile(fpath, &f); err != nil {
			return
		}

	case ".json":
		if err = decodeJSONFile(fpath, &f); err != nil {
			return
		}

	default:
		err = fmt.Errorf("unsupported extension provided: %s", ext)
	}

	// cron.New(f.sendAllNotifications).Every(time.Minute)
	fp = &f
	return
}

// Figma is a base structure
type Figma struct {
	hc http.Client
	// out *journaler.Journaler

	AccessToken string `json:"accesToken" toml:"accessToken"`
}

// request is a base request wrapper
func (f *Figma) request(m string, u *url.URL, b io.Reader) (response *http.Response, err error) {
	var req *http.Request
	if req, err = http.NewRequest(m, u.String(), b); err != nil {
		return
	}

	req.Header.Add("X-Figma-Token", f.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	return f.hc.Do(req)
}

func (f *Figma) getFile(fileKey string) (entry *EntryFile, err error) {
	var u *url.URL
	uPath := fmt.Sprintf(filePath, fileKey)
	if u, err = newURL(baseURL, uPath); err != nil {
		return
	}

	var response *http.Response
	if response, err = f.request(http.MethodGet, u, nil); err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = handleFileErr(response)
		return
	}

	if err = json.NewDecoder(response.Body).Decode(&entry); err != nil {
		return
	}

	if err = getDefaults(entry.Document); err != nil {
		return
	}

	return
}

func (f *Figma) getNodes(fileKey string, nodeIDs []string) (entry *EntryNodes, err error) {
	var u *url.URL
	uPath := fmt.Sprintf(nodesPath, fileKey)
	if u, err = newURL(baseURL, uPath); err != nil {
		return
	}

	var searchParams = url.Values{}
	searchParams.Add("ids", strings.Join(nodeIDs, ","))
	u.RawQuery = searchParams.Encode()

	var response *http.Response
	if response, err = f.request(http.MethodGet, u, nil); err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = handleFileErr(response)
		return
	}

	if err = json.NewDecoder(response.Body).Decode(&entry); err != nil {
		return
	}

	return
}

func (f *Figma) getStylesByTeamID(teamID string) (colors []*TeamColor, gradients []*TeamGradient, texts []*TeamText, effects []*TeamEffect, err error) {
	var u *url.URL
	uPath := fmt.Sprintf(stylesPath, teamID)
	if u, err = newURL(baseURL, uPath); err != nil {
		return
	}

	var response *http.Response
	if response, err = f.request(http.MethodGet, u, nil); err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = handleStyleResponseErr(response)
		return
	}

	var styles EntryStyles
	if err = json.NewDecoder(response.Body).Decode(&styles); err != nil {
		return
	}

	m := make(map[string][]*NodeStyleType)

	for _, v := range styles.Meta.Styles {
		m[v.FileKey] = append(m[v.FileKey], newNodeStyleType(v.NodeID, v.StyleType))
	}

	for fileKey, node := range m {
		n := make(map[StyleType][]string)

		for _, ns := range node {
			if err = getStyleTypesMap(n, ns); err != nil {
				return
			}
		}

		if colors, gradients, texts, effects, err = setStyleTypes(f, fileKey, n); err != nil {
			return
		}
	}

	return
}

// GetFile returns a figma file as a node entry
func (f *Figma) GetFile(fileKey string) (entry *EntryFile, err error) {
	return f.getFile(fileKey)
}

// GetNodes returns a node metadata entry containing a map of nodes keyed by node ID
func (f *Figma) GetNodes(filekey string, nodeIDs []string) (entry *EntryNodes, err error) {
	return f.getNodes(filekey, nodeIDs)
}

// GetStylesByTeamID returns a team's colors as a map keyed by color name
func (f *Figma) GetStylesByTeamID(teamID string) (colors []*TeamColor, gradients []*TeamGradient, texts []*TeamText, effects []*TeamEffect, err error) {
	return f.getStylesByTeamID(teamID)
}
