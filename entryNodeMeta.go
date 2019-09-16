package figma

// EntryNodeMeta is the metadata of a GET File Nodes response from Figma API
//  NOTE: contains the node entry
type EntryNodeMeta struct {
	// Document represents the document entry for a node
	Document EntryNode `json:"document"`
	// Components represent the components that utilize a node (not used in this libary)
	Components interface{} `json:"components"`
	// SchemaVersion represents the schema verison of a node
	SchemaVersion int `json:"schemaVersion"`
	// Styles represents the styles that utilize a node (not used in this libraary)
	Styles interface{} `json:"styles"`
}
