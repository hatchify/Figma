package figma

// Pattern type
const (
	Columns PatternType = "COLUMNS" // Vertical grid
	Rows    PatternType = "ROWS"    // Horizontal grid
	Grid    PatternType = "GRID"    // Square grid
)

// Alignment type
const (
	Min     AlignmentType = "MIN"     // Grid starts at the left or top of the frame
	Stretch AlignmentType = "STRETCH" // Grid is stretched to fit the frame
	ACenter AlignmentType = "CENTER"  // Grid is center aligned
)

// EntryLayoutGrid represents the layout grid entry
type EntryLayoutGrid struct {
	// Pattern is the orientation of the grid as a PatterType
	Pattern PatternType `json:"pattern"`
	// SectionSize is the width of column grid or height of row grid or square grid spacing
	SectionSize float64 `json:"sectionSize"`
	// Visible is visibility bool
	Visible bool `json:"Visible"`
	// Color is the color of the grid
	Color *EntryColor `json:"color"`

	// The following properties are only meaningful for directional grids (COLUMNS or ROWS)
	// Aligment is the position of a grid as AlignmentType
	Alignment AlignmentType `json:"alignment,omitempty"`
	// Gutter size is the spacing in between columns and rows
	GutterSize float64 `json:"gutterSize,omitempty"`
	// Offset is the spacing before the first column or row
	Offset float64 `json:"offset,omitempty"`
	// Count is the number of columns or rows
	Count float64 `json:"count,omitempty"`
}

// PatternType represents the various pattern types
type PatternType string

// AlignmentType represents the various aligment types
type AlignmentType string
