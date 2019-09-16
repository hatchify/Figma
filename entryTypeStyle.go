package figma

// Text case type
const (
	Original TextCaseType = "ORIGINAL"
	Upper    TextCaseType = "UPPER"
	Lower    TextCaseType = "LOWER"
	Title    TextCaseType = "TITLE"
)

// Text decoration type
const (
	NoDecoration  TextDecorationType = "NONE"
	Strikethrough TextDecorationType = "STRIKETHROUGH"
	Underline     TextDecorationType = "UNDERLINE"
)

// Text align horizontal type
const (
	TextLeft      TextAlignHorizontalType = "LEFT"
	TextRight     TextAlignHorizontalType = "RIGHT"
	TextHCenter   TextAlignHorizontalType = "CENTER"
	TextJustified TextAlignHorizontalType = "JUSTIFIED"
)

// Text align vertical type
const (
	TextTop     TextAlignVerticalType = "TOP"
	TextVCenter TextAlignVerticalType = "CENTER"
	TextBottom  TextAlignVerticalType = "BOTTOM"
)

// Line height unit
const (
	Pixels LineHeightUnit = "PIXELS"

	FontSizePercentage  LineHeightUnit = "FONT_SIZE_%"
	IntrinsicPercentage LineHeightUnit = "INTRINSIC_%"
)

// EntryTypeStyle is the type style entry
type EntryTypeStyle struct {
	// FontFamily is the font family of a TEXT style's font
	FontFamily string `json:"fontFamily"`
	// FontPostScriptName is the post script name of a font
	FontPostScriptName string `json:"fontPostScriptName"`
	// ParagraphSpacing is the space between paragraphs in px
	ParagraphSpacing float64 `json:"paragraphSpacing"`
	// ParagraphIndent is the paragraph indentation
	ParagraphIndent float64 `json:"paragraphIndent"`
	// Italic determines whether or not text is italicized
	Italic bool `json:"italic"`
	// FontWeight is the weight of a font
	FontWeight float64 `json:"fontWeight"`
	// FontSize is the size of a font
	FontSize float64 `json:"fontSize"`
	// TextCase is the text casing applied to the node
	//  NOTE: default does not return, default value is Original...
	TextCase *TextCaseType `json:"textCase,omitempty"`
	// TextDecoration is the text decoration applied to the node
	//  NOTE: default does not return, default value is NoDecoration
	TextDecoration *TextDecorationType `json:"textDecoration,omitempty"`
	// TextAlignHorizontal is the horizontal alignment of a TEXT style
	TextAlignHorizontal TextAlignHorizontalType `json:"textAlignHorizontal"`
	// TextAlignVertical is the vertical alignment of a TEXT style
	TextAlignVertical TextAlignVerticalType `json:"textAlignVertical"`
	// LetterSpacing is the letter spacing of a TEXT style
	LetterSpacing float64 `json:"letterSpacing"`
	// Fills is the slice of paints applied to characters
	Fills []EntryPaint `json:"fills"`
	// LineHeightPX is the height of a TEXT style in px
	LineHeightPX float64 `json:"lineHeightPx"`
	// LineHeightPercent is the height of a TEXT style in percentage points
	//  NOTE: default does not return, default value is 100.0...
	LineHeightPercent *float64 `json:"lineHeightPercent,omitempty"`
	// LineHeightPercentFontSize is the height of a TEXT style in percentag points relative to the style's font size
	LineHeightPercentFontSize float64 `json:"lineHeightPercentFontSize"`
	// LineHeightUnit is the unit measurement for the line height of a TEXT style
	LineHeightUnit LineHeightUnit `json:"lineHeightUnit"`
}

// GetDefaults is a helper method for the json.Decoding process
func (e *EntryTypeStyle) getDefaults() (err error) {
	ec := *e
	if v := Original; e.TextCase == nil {
		ec.TextCase = &v
	} else {
		ec.TextCase = e.TextCase
	}

	if v := NoDecoration; e.TextDecoration == nil {
		ec.TextDecoration = &v
	} else {
		ec.TextDecoration = e.TextDecoration
	}

	if v := 100.0; e.LineHeightPercent == nil {
		ec.LineHeightPercent = &v
	} else {
		ec.LineHeightPercent = e.LineHeightPercent
	}

	*e = ec
	return
}

// TextCaseType represents the various text case types
type TextCaseType string

// TextDecorationType represents the various text decoration types
type TextDecorationType string

// TextAlignHorizontalType represents the various text align horizontal types
type TextAlignHorizontalType string

// TextAlignVerticalType represents the various text align vertical types
type TextAlignVerticalType string

// LineHeightUnit represents the various line height units
type LineHeightUnit string
