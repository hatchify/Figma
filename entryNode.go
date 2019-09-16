package figma

// Node type
const (
	Line  NodeType = "LINE"
	Text  NodeType = "TEXT"
	Star  NodeType = "STAR"
	Frame NodeType = "FRAME"
	Group NodeType = "GROUP"
	Slice NodeType = "SLICE"

	Vector    NodeType = "VECTOR"
	Canvas    NodeType = "CANVAS"
	Ellipse   NodeType = "ELLIPSE"
	Document  NodeType = "DOCUMENT"
	Instance  NodeType = "INSTANCE"
	Rectangle NodeType = "RECTANGLE"
	Component NodeType = "COMPONENT"

	RegularPolygon   NodeType = "REGULAR_POLYGON"
	BooleanOperation NodeType = "BOOLEAN_OPERATION"
)

// Normal blend type
const (
	PassThrough BlendModeType = "PASS_THROUGH" // Only applicable to objects with children
	Normal      BlendModeType = "NORMAL"
)

// Darken blend type
const (
	Darken     BlendModeType = "DARKEN"
	Multiply   BlendModeType = "MULTIPLY"
	LinearBurn BlendModeType = "LINEAR_BURN"
	ColorBurn  BlendModeType = "COLOR_BURN"
)

// Lighten blend type
const (
	Lighten     BlendModeType = "LIGHTEN"
	Screen      BlendModeType = "SCREEN"
	LinearDodge BlendModeType = "LINEAR_DODGE"
	ColorDodge  BlendModeType = "COLOR_DODGE"
)

// Contrast blend type
const (
	Overlay   BlendModeType = "OVERLAY"
	SoftLight BlendModeType = "SOFT_LIGHT"
	HardLight BlendModeType = "HARD_LIGHT"
)

// Inversion blend type
const (
	Difference BlendModeType = "DIFFERENCE"
	Excusion   BlendModeType = "EXCLUSION"
)

// Component blend type
const (
	Hue BlendModeType = "HUE"

	Saturation     BlendModeType = "SATURATION"
	ComponentColor BlendModeType = "COLOR"
	Luminosity     BlendModeType = "LUMINOSITY"
)

// Easing Type
const (
	EaseIn  EasingType = "EASE_IN"  // Ease in with an animation curve similar to CSS ease-in
	EaseOut EasingType = "EASE_OUT" // Ease out with an animation curve similar to CSS ease-out

	EaseInAndOut EasingType = "EASE_IN_AND_OUT" // Ease in and then out with an animation curve similar to CSS ease-in-out
)

// StrokeCap type
const (
	NoStrokeCap StrokeCapType = "NONE"

	CRound StrokeCapType = "ROUND"
	Square StrokeCapType = "SQUARE"

	LineArrow     StrokeCapType = "LINE_ARROW"
	TriangleArrow StrokeCapType = "TRIANGLE_ARROW"
)

// StrokeJoin type
const (
	Miter  StrokeJoinType = "MITER"
	Bevel  StrokeJoinType = "BEVEL"
	JRound StrokeJoinType = "ROUND"
)

// StrokeAlign type
const (
	Inside  StrokeAlignType = "INSIDE"  // draw stroke inside the shape boundary
	Outside StrokeAlignType = "OUTSIDE" // draw stroke outside the shape boundary
	Center  StrokeAlignType = "CENTER"  // draw stroke centered along the shape boundary
)

// BooleanOperation type
const (
	Union     BooleanOperationType = "UNION"
	Intersect BooleanOperationType = "INTERSECT"
	Subtract  BooleanOperationType = "SUBTRACT"
	Exclude   BooleanOperationType = "EXCLUDE"
)

// EntryNode represents the node entry
type EntryNode struct {
	// Below are the fields that are core to every node entry
	// ID represents the id of a node
	ID string `json:"id"`
	// Name represents the name of a node's document
	Name string `json:"name"`
	// Type represents the node's type
	Type NodeType `json:"Type"`
	// Visible determines whether or not the node is visible on the canvas
	//  NOTE: default does not return, default value is true...
	Visible *bool `json:"visible,omitempty"`

	// Below are the fields that are dependent on the node type returned
	//  NOTE: in alphabetical order for simplicity
	//  NOTE: some fields overlap different types... go to www.figma.com/developers/api and search "Node Types" for more info

	// AbsoluteBoundingBox represents the bounds of a node in Figma (not used in this library)
	AbsoluteBoundingBox interface{} `json:"absoluteBoundingBox,omitempty"`
	// Background is the background of a node
	Background []EntryPaint `json:"background,omitempty"`
	// BackgroundColor is the background color of the canvase
	BackgroundColor *EntryColor `json:"backgroundColor,omitempty"`
	// Blend mode represents the blend mode of a node
	BlendMode BlendModeType `json:"blendMode,omitempty"`
	// BooleanOperation indicates the type of boolean operation applied as a BooleanOperationType
	BooleanOperation BooleanOperationType `json:"booleanOperation,omitempty"`
	// Characters represents the character type assigned to a node
	Characters string `json:"characters,omitempty"`
	// CharacterStyleOverrides are the overrides applied to a node
	CharacterStyleOverrides []float64 `json:"characterStyleOverrides,omitempty"`
	// Children is an array of nodes attached to this node
	Children []EntryNode `json:"children,omitempty"`
	// ClipsContent determines whether or not this node clips content outside of its bounds
	ClipsContent bool `json:"clipsContent,omitempty"`
	// ComponentID is the id of the compenent that this instance cam from, refers to components table
	ComponentID string `json:"componentId,omitempty"`
	// Constraints represents the contraints of a node in Figma
	Constraints EntryLayoutConstraint `json:"constraints,omitempty"`
	// Effects represents the effects assigned to a node
	Effects []EntryEffect `json:"effects,omitempty"`
	// ExportSettings is the slice of export settings representing images to export from the canvas
	ExportSettings []EntryExportSetting `json:"exportSettings,omitempty"`
	// Fills represet the fill settings of a node
	Fills []EntryPaint `json:"fills,omitempty"`
	// FillGeometry is not used in this libary
	FillGeometry interface{} `json:"fillGeometry,omitempty"`
	// IsMask determines whether a node masks sibling nodes in front of it
	IsMask bool `json:"isMask,omitempty"`
	// IsMaskOutline dtermines wheter or not a mask ignores fill style and effects
	IsMaskOutline bool `json:"isMaskOutline,omitempty"`
	// LayoutGrids is a slice of layout grids attached to this node
	LayoutGrids []EntryLayoutGrid `json:"layoutGrids,omitempty"`
	// Locked determines if a layer is locked and cannot be edited
	Locked bool `json:"locked,omitempty"`
	// Opacity is the opacity of the node
	//  NOTE: default does not return, default value is 1.0...
	Opacity *float64 `json:"opacity,omitempty"`
	// PreserveRatio determines whether to keep height and width constrained to same ratio
	PreserveRatio bool `json:"preserveRatio,omitempty"`
	// PrototypeStartNodeID is the node ID that corresponds to the start frame for prototypes
	PrototypeStartNodeID string `json:"prototypeStartNodeID,omitempty"`
	// RelativeTransform is the top two rows of a matrix that represents the 2D transform of this node relative to its parent
	//  NOTE: Only present if geometry=paths query param is passed
	RelativeTransform [][]float64 `json:"relativeTransform,omitempty"`
	// Size is the width/height of an element
	//  NOTE: Only present if geometry=paths query param is passed
	Size EntryVector `json:"size,omitempty"`
	// StrokeAlign represents the alignment of the strokes assigned to a node
	StrokeAlign string `json:"strokeAlign,omitempty"`
	// StrokeCap describes the  end caps of vector paths as a StrokeCapType
	StrokeCap StrokeCapType `json:"strokeCap,omitempty"`
	// StrokeDashes is a slice of floating point numbers describing the pattern of dash length and gap lengths that the vector path follows
	StrokeDashes []float64 `json:"strokeDashes,omitempty"`
	// StrokeGeometry is a slice of paths representing the object stroke (not used in this libary)
	//  NOTE: no type information provided for stroke geometry on Figma API documentation.  Supposed to be a slice of type Path ([]Path)
	//  NOTE: Only present if geometry=paths query param is passed
	StrokeGeometry interface{}
	// StrokeJoin describes how the cornes in vector paths are rendered
	StrokeJoin StrokeJoinType `json:"strokeJoin,omitempty"`
	// StrokeMiterAngle is the corner angle, in degreees, below which StrokeJoin will be set to "BEVEL" to avoid super sharp corners
	//  NOTE: default does not return, default value is 28.96...
	StrokeMiterAngle *float64 `json:"strokeMiterAngle,omitempty"`
	// Strokes represents the strokes associated with a node
	Strokes []EntryPaint `json:"strokes,omitempty"`
	// StrokeWeight represents the stroke weight applied to a node
	StrokeWeight float64 `json:"strokeWeight,omitempty"`
	// StyleOverrideTable is the table of overrides of a style
	StyleOverrideTable map[int]EntryTypeStyle `json:"styleOverrideTable,omitempty"`
	// Style is the style of text including font family and weight
	Style *EntryTypeStyle `json:"style,omitempty"`
	// Styles is a mapping of StyleType to style ID of styles present on this node
	Styles interface{} `json:"styles,omitempty"`
	// TransitionDuration is the duration of the prototyping transition on this node (in milliseconds)
	TransitionDuration float64 `json:"transitionDuration,omitempty"`
	// TransitionEasing is the easing curve used in the prototyping transition on this node
	TransitionEasing EasingType `json:"transitionEasing"`
	// TransitionNodeID is the node ID of node to transition to in prototyping
	TransitionNodeID string `json:"transitionNodeID,omitempty"`
}

// GetDefaults is a helper method for the json.Decoding process
func (e *EntryNode) getDefaults() (err error) {
	ec := *e
	if v := true; e.Visible == nil {
		ec.Visible = &v
	} else {
		ec.Visible = e.Visible
	}

	if v := 1.0; e.Opacity == nil {
		switch e.Type {
		case Slice:
		default:
			ec.Opacity = &v
		}
	} else {
		ec.Opacity = e.Opacity
	}

	if v := 28.96; e.StrokeMiterAngle == nil {
		ec.StrokeMiterAngle = &v
	} else {
		ec.StrokeMiterAngle = e.StrokeMiterAngle
	}

	*e = ec
	return
}

// NodeType represents the various node types
type NodeType string

// BlendModeType represents the various blend mode types
type BlendModeType string

// EasingType represents the various easing types
type EasingType string

// StrokeCapType represents the various stroke cap types
type StrokeCapType string

// StrokeJoinType represents the various stroke join types
type StrokeJoinType string

// StrokeAlignType represents the various stroke align types
type StrokeAlignType string

// BooleanOperationType represents the various boolean operation types
type BooleanOperationType string
