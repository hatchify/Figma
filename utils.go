package figma

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func decodeJSONFile(fname string, value interface{}) (err error) {
	var f *os.File
	if f, err = os.Open(fname); err != nil {
		return
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(value)
}

func newURL(base, path string) (u *url.URL, err error) {
	if u, err = url.Parse(base); err != nil {
		return
	}

	u.Path = path
	return
}

func decodeJSONResponse(response *http.Response, v interface{}) (err error) {
	if err = json.NewDecoder(response.Body).Decode(&v); err != nil {
		return
	}

	return
}

func handleFileErr(response *http.Response) (err error) {
	var entry EntryNodes
	if err = json.NewDecoder(response.Body).Decode(&entry); err != nil {
		err = errors.New(response.Status)
		return
	}

	if entry.Err != "" {
		err = errors.New(entry.Err)
	} else {
		err = errors.New("unnexpected error")
	}

	return
}

func handleStyleResponseErr(response *http.Response) (err error) {
	var styles EntryStyles
	if err = json.NewDecoder(response.Body).Decode(&styles); err != nil {
		err = errors.New(response.Status)
		return
	}

	if styles.Message != "" {
		err = errors.New(styles.Message)
	} else if styles.Err != "" {
		err = errors.New(styles.Err)
	} else {
		err = errors.New("unnexpected error")
	}

	return
}

func getStyleTypesMap(n map[StyleType][]string, ns *NodeStyleType) (err error) {
	switch ns.StyleType {
	case StyleFill:
		n[StyleFill] = append(n[StyleFill], ns.NodeID)
	case StyleText:
		n[StyleText] = append(n[StyleText], ns.NodeID)
	case StyleEffect:
		n[StyleEffect] = append(n[StyleEffect], ns.NodeID)
	default:
		err = fmt.Errorf("invalid type: exepcting one of \"%s\", \"%s\", or \"%s\", received: %s", StyleFill, StyleText, StyleEffect, ns.StyleType)
		return
	}

	return
}

func setStyleTypes(f *Figma, fileKey string, n map[StyleType][]string) (colors []*TeamColor, gradients []*TeamGradient, texts []*TeamText, effects []*TeamEffect, err error) {
	var entry *EntryNodes

	for styleType, nodeIDs := range n {
		switch styleType {
		case StyleFill:
			if entry, err = f.getNodes(fileKey, nodeIDs); err != nil {
				return
			}

			for _, node := range entry.Nodes {
				if err = getDefaults(&node.Document); err != nil {
					return
				}

				name := node.Document.Name
				for k, fill := range node.Document.Fills {
					if err = getDefaults(&fill); err != nil {
						return
					}

					if k > 1 {
						err = errors.New("unable to assign more than one fill to a fill entry")
						return
					}

					t := fill.Type
					if t == Solid {
						c := newTeamColor(name, fill.Color)
						colors = append(colors, c)
					}

					if t == GradientAngular || t == GradientDiamond || t == GradientLinear || t == GradientRadial {
						g := newTeamGradient(name, fill.BlendMode, &fill.GradientHandlePositions, &fill.GradientStops)
						gradients = append(gradients, g)
					}
				}

			}
		case StyleText:
			if entry, err = f.getNodes(fileKey, nodeIDs); err != nil {
				return
			}

			for _, node := range entry.Nodes {
				if err = getDefaults(&node.Document); err != nil {
					return
				}

				if err = getDefaults(node.Document.Style); err != nil {
					return
				}

				name := node.Document.Name

				t := newTeamText(name, node.Document.Characters, node.Document.Style)
				texts = append(texts, t)
			}
		case StyleEffect:
			if entry, err = f.getNodes(fileKey, nodeIDs); err != nil {
				return
			}

			for _, node := range entry.Nodes {
				if err = getDefaults(&node.Document); err != nil {
					return
				}

				name := node.Document.Name
				for k, effect := range node.Document.Effects {
					if k > 1 {
						err = errors.New("unable to assign more than one effect to an effect entry")
						return
					}

					e := newTeamEffect(name, &effect)
					effects = append(effects, e)
				}

			}
		}
	}

	return
}

// Defaults is an interface for calling method GetDefaults
type Defaults interface {
	getDefaults() (err error)
}

func getDefaults(entry Defaults) (err error) {
	if err = entry.getDefaults(); err != nil {
		return
	}

	return
}
