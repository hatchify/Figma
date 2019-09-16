package figma

func newTeamColor(n string, ec *EntryColor) (tcp *TeamColor) {
	var tc TeamColor
	tc.Name = n
	tc.Color = ec
	return &tc
}

// TeamColor is the data structure for a solid fill
type TeamColor struct {
	Name  string
	Color *EntryColor
}
