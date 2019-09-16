package figma

// Constraint type
const (
	Scale  ConstraintType = "SCALE"  // Scale by value
	Width  ConstraintType = "WIDTH"  // Scale proportionally and set width to value
	Height ConstraintType = "HEIGHT" // Scale proportionally and set height to value
)

// EntryConstraint represents the constraint entry
type EntryConstraint struct {
	// Type is the type of constraint to apply
	Type ConstraintType `json:"type"`
	// Value is the value applied to the ConstraintType
	Value float64 `json:"value"`
}

// ConstraintType represnets the various constraint types
type ConstraintType string
