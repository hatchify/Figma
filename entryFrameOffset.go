package figma

// EntryFrameOffset is the frame offset entry
type EntryFrameOffset struct {
	// NodeID is the node id
	NodeID string `json:"node_id"`
	// NodeOffset is the 2d vector offset within the frame
	NodeOffset EntryVector `json:"node_offset"`
}
