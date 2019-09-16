package figma

func newNodeStyleType(nodeID string, styleType StyleType) (np *NodeStyleType) {
	var n NodeStyleType
	n.NodeID = nodeID
	n.StyleType = styleType
	return &n
}

// NodeStyleType is an internal architecture
type NodeStyleType struct {
	// NodeID is the id of a node
	NodeID string
	// StyleType is the style type associated with a node
	StyleType StyleType
}
