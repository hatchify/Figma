package figma

func newTeamEffect(n string, e *EntryEffect) (tep *TeamEffect) {
	var te TeamEffect
	te.Name = n
	te.Effect = e
	return &te
}

// TeamEffect is the data structure for an effect
type TeamEffect struct {
	Name   string
	Effect *EntryEffect
}
