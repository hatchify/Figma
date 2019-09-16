package figma

// EntryStyleMeta contains the metadata of a GET Styles response from Figma API
//  NOTE: contains the style entry
type EntryStyleMeta struct {
	// Cursor represents a map (map[string]int) that indicates the starting/ending point from which objects are returned (not used in this library)
	Cursor interface{} `json:"cursor"`
	// Styles represents a slice of style entries
	Styles []EntryStyle `json:"styles"`
}
