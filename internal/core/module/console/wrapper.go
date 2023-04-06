package console

type Wrapper struct {
	*App
}

func New() *Wrapper {
	app := &App{}
	return &Wrapper{
		app,
	}
}
