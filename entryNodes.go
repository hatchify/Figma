package figma

// EntryNodes is the top level data structure for a GET file nodes response from Figma API
type EntryNodes struct {
	// Status represents the response code of a File response
	//  NOTE: this is only given when an err is thrown
	Status int `json:"status,omitempty"`
	// Err represents the error message if request fails
	Err string `json:"err,omitempty"`
	// Name represents the name of a file
	Name string `json:"name,omitemtpy"`
	// LastModified is the string of the last modified date-time in RFC3339 format
	LastModified string `json:"lastModified,omitempty"`
	// ThumbnailURL represents the thumbnail url for a file
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`
	// Version represents a file's version
	Version string `json:"version,omitempty"`
	// Nodes represents a map of nodes
	Nodes map[string]EntryNodeMeta `json:"nodes,omitempty"`
}
