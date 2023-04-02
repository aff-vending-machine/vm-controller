package property

type Screen struct {
	Rotate int
	Width  int
	Height int
}

func NewScreen(r, w, h int) *Screen {
	return &Screen{r, w, h}
}
