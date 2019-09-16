package figma

func newTeamGradient(n string, bm BlendModeType, ghp *[]EntryVector, gs *[]EntryColorStop) (tgp *TeamGradient) {
	var tg TeamGradient
	tg.Name = n
	tg.BlendMode = bm
	tg.GradientStops = gs
	tg.GradientHandlePositions = ghp
	return &tg
}

// TeamGradient is the data structure for a solid fill
type TeamGradient struct {
	Name      string
	BlendMode BlendModeType

	GradientStops *[]EntryColorStop

	GradientHandlePositions *[]EntryVector
}
