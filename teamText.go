package figma

func newTeamText(n, c string, et *EntryTypeStyle) (ttp *TeamText) {
	var tt TeamText
	tt.Name = n
	tt.Characters = c
	tt.Style = et
	return &tt
}

// TeamText is the data structure for a solid fill
type TeamText struct {
	Name  string
	Style *EntryTypeStyle

	Characters string
}
