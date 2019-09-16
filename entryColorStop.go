package figma

// EntryColorStop is a gradient stop entry
type EntryColorStop struct {
	// Color represents the color of a gradient stop in RGBA
	Color *EntryColor `json:"color"`
	// Position is the position of a gradient stop
	Position float64 `json:"position"`
}
