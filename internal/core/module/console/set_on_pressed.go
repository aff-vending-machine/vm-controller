package console

func (app *App) SetOnPressed(handler Handler) {
	app.handler = &handler
}
