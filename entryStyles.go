package figma

// EntryStyles represents the StyleResponse entry
type EntryStyles struct {
	// Error represents if error is present
	Error bool `json:"error,omitempty"`
	// Status represent the status code of the response
	Status int `json:"status"`
	// Message represent the error message if error is present
	Message string `json:"message,omitempty"`
	// Err represents the error message if request fails and error bool is not returned
	Err string `json:"err,omitempty"`
	// Meta represents the metadata of the response
	Meta EntryStyleMeta `json:"meta,omitempty"`
}
