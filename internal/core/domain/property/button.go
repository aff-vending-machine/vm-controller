package property

type Button struct {
	Label  string
	PosX   int
	PosY   int
	Width  int
	Height int
}

func NewLeftButton(label string) *Button {
	return &Button{
		Label:  label,
		PosX:   140,
		PosY:   2250,
		Width:  400,
		Height: 200,
	}
}

func NewRightButton(label string) *Button {
	return &Button{
		Label:  label,
		PosX:   900,
		PosY:   2250,
		Width:  400,
		Height: 200,
	}
}
