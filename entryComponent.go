package figma

// EntryComponent is the component entry
type EntryComponent struct {
	// Key  is the unique identifier of the component
	Key string `json:"key"`
	// FileKey is the unique identifier of the figma file which contains the component
	FileKey string `json:"file_key,omitempty"`
	// NodeID is the id of the component node within the figma file
	NodeID string `json:"node_id,omitempty"`
	// ThumbnailURL is the URL link to the component's thumbnail image
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Name is the name of the component
	Name string `json:"name"`
	// Description is the description of the component
	Description string `json:"description"`
	// CreatedAt is the string of the created at date-time in RFC3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the string of the updated at date-time in RFC3339 format
	UpdatedAt string `json:"updated_at,omitempty"`
	// User is the user entry of the user that last updated the component
	User EntryUser `json:"user,omitempty"`
	// ContainingFrame is the data on component's containing frame, if component resides within a frame
	ContainingFrame EntryFrameInfo `json:"containing_frame,omitempty"`
}
