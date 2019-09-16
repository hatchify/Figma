package figma

// Style type
const (
	StyleFill   StyleType = "FILL"
	StyleText   StyleType = "TEXT"
	StyleEffect StyleType = "EFFECT"
	StyleGrid   StyleType = "GRID"
)

// EntryStyle represents the Style entry
type EntryStyle struct {
	// Key is the key of the style
	Key string `json:"key"`
	// FileKey represents the file that owns the style
	FileKey string `json:"file_key,omitempty"`
	// NodeID is the node ID of a style
	NodeID string `json:"node_id,omitempty"`
	// StyleType represents the style type of a style
	StyleType StyleType `json:"style_type"`
	// ThumbnailURL is the thumbnail url for a style
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Name is the name of a style
	Name string `json:"name"`
	// Description is the description of a style
	Description string `json:"description"`
	// CreatedAt is the string of the created at date-time in RFC3339 format
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt is the string of the updated at date-time in RFC3339 format
	UpdatedAt string `json:"updated_at,omitempty"`
	// User is the user entry of the user that last updated the style
	User EntryUser `json:"user,omitempty"`
	// SortPosition represents the sort position of a style
	SortPosition string `json:"sort_position,omitempty"`
}

// StyleType represents the various style types
type StyleType string
