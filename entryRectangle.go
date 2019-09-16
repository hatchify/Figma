package figma

// EntryRectangle represents the rectangle entry
type EntryRectangle struct {
	// X is the xcoordinate of top left corner of the rectangle
	X float64 `json:"x"`
	// Y is the y coordinate of top left corner of the rectangle
	Y float64 `json:"y"`
	// Width is the width of a rectangle
	Width float64 `json:"width"`
	// Height is the height of a rectangle
	Height float64 `json:"height"`
}
