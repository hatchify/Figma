package figma

// EntryFrameInfo is the frame info entry
type EntryFrameInfo struct {
	// NodeID is the id of the frame node within the figma file
	NodeID string `json:"node_id"`
	// Name is the name of the frame
	Name string `json:"name"`
	// BackgroundColor is the background color of the frame
	BackgroundColor string `json:"background_color"`
	// PageID is the id of the frame's residing page
	PageID string `json:"page_id"`
	// PageName is the name of the frame's residing page
	PageName string `json:"page_name"`
}
