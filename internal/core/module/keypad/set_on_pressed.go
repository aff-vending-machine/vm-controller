package keypad

func (app *App) SetOnPressed(handler Handler) {
	app.handler = &handler
}
