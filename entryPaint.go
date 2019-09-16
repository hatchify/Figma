package figma

// Paint type
const (
	Solid PaintType = "SOLID"
	Image PaintType = "IMAGE"
	Emoji PaintType = "EMOJI"

	GradientLinear  PaintType = "GRADIENT_LINEAR"
	GradientRadial  PaintType = "GRADIENT_RADIAL"
	GradientAngular PaintType = "GRADIENT_ANGULAR"
	GradientDiamond PaintType = "GRADIENT_DIAMOND"
)

// Scale mode type
const (
	Fill ScaleModeType = "FILL"
	Fit  ScaleModeType = "FIT"
	Tile ScaleModeType = "TILE"

	SStretch ScaleModeType = "STRETCH"
)

// EntryPaint represents a Fill entry
type EntryPaint struct {
	// Type represents the type of a fill
	Type PaintType `json:"type"`
	// Visible determines whether or not a paint is visible
	//  NOTE: default does not return, default value is true...
	Visible *bool `json:"visible,omitempty"`
	// Opacity is the overall opacity of the paint
	//  NOTE: default does not return, default value is 1.0...
	Opacity *float64 `json:"opacity,omitempty"`

	// NOTE: below entries are only returned for a solid paint
	// Color represents the color of a fill in RGBA
	Color *EntryColor `json:"color,omitempty"`

	// NOTE: below entries are only returned for a gradient paint
	// BlendMode is How this node blends with nodes behind it in the scene
	BlendMode BlendModeType `json:"blendMode,omitempty"`
	// GradientHandlePositions are positioning coordinates for a gradient
	GradientHandlePositions []EntryVector `json:"gradientHandlePositions,omitempty"`
	// GradientStops are the stops for a gradient
	GradientStops []EntryColorStop `json:"gradientStops,omitempty"`

	// NOTE: below entries are only returned for an image paint
	// ScaleMode is the imaging scaling mode as ScaleModeType
	ScaleMode ScaleModeType `json:"scaleMode,omitempty"`
	// ImageTransform is the affine transformation matrix applied to the image, only present if scaleMode is STRETCH
	ImageTransform [][]float64 `json:"transform,omitempty"`
	// ScalingFactor is the amount an image is scaled by in tiling, only present if scaleMode is TILE
	ScalingFactor float64 `json:"scalingFactor,omitempty"`
	// ImageRef is a reference to an image embedded in the file
	ImageRef string `json:"imageRef,omitempty"`
}

// GetDefaults is a helper method for the json.Decoding process
func (e *EntryPaint) getDefaults() (err error) {
	ec := *e
	if v := true; e.Visible == nil {
		ec.Visible = &v
	} else {
		ec.Visible = e.Visible
	}

	if v := 1.0; e.Opacity == nil {
		ec.Opacity = &v
	} else {
		ec.Opacity = e.Opacity
	}

	*e = ec
	return
}

// PaintType represents the various paint types
type PaintType string

// ScaleModeType represents the various scale mode types
type ScaleModeType string
