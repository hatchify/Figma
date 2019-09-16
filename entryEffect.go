package figma

// Effect type
const (
	InnerShadow    EffectType = "INNER_SHADOW"
	DropShadow     EffectType = "DROP_SHADOW"
	LayerBlur      EffectType = "LAYER_BLUR"
	BackgroundBlur EffectType = "BACKGROUND_BLUR"
)

// EntryEffect represents the effect entry of a style
type EntryEffect struct {
	// Type represents the type of effect
	Type string `json:"type"`
	// Visible represents the visibility of an effect
	Visible bool `json:"visible"`
	// Radius is the radius of the blur effect
	Radius float64 `json:"radius,omitempty"`

	// The following properties are for shadows only
	// Color represents the color of a fill in RGBA
	Color *EntryColor `json:"color,omitempty"`
	// BlendMode is the blend mode of an effect
	BlendMode BlendModeType `json:"blendMode,omitempty"`
	// Offset is the offset of an effect
	Offset EntryVector `json:"offset,omitempty"`
}

// EffectType represents the various effect types
type EffectType string
