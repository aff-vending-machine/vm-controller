package keypad

type InputKey interface {
	OnPressed(key string) error
}
