package figma

// Vertical constraint type
const (
	Top       VerticalConstraintType = "TOP"        // Node is laid out relative to top of the containing frame
	VScale    VerticalConstraintType = "SCALE"      // Node scales vertically with containing frame
	Bottom    VerticalConstraintType = "BOTTOM"     // Node is laid out relative to bottom of the containing frame
	VCenter   VerticalConstraintType = "CENTER"     // Node is vertically centered relative to containing frame
	TopBottom VerticalConstraintType = "TOP_BOTTOM" // Both top and bottom of node are constrained relative to containing frame (node stretches with frame)
)

// Horizontal constraint type
const (
	Left      HorizontalConstraintType = "LEFT"       // Node is laid out relative to left of the containing frame
	Right     HorizontalConstraintType = "RIGHT"      // Node is laid out relative to right of the containing frame
	HCenter   HorizontalConstraintType = "CENTER"     // Node is horizontally centered relative to containing frame
	LeftRight HorizontalConstraintType = "LEFT_RIGHT" // Both left and right of node are constrained relative to containing frame (node stretches with frame)
	HScale    HorizontalConstraintType = "SCALE"      // Node scales horizontally with containing frame
)

// EntryLayoutConstraint is the layout constraint entry
type EntryLayoutConstraint struct {
	// Vertical is a vertical constraint
	Vertical VerticalConstraintType `json:"vertical"`
	// Horizontal is a horizontal constraint
	Horizontal HorizontalConstraintType `json:"horizontal"`
}

// VerticalConstraintType represents the various vertical constraint types
type VerticalConstraintType string

// HorizontalConstraintType represents the various horizontal constraint types
type HorizontalConstraintType string
