package request

type Get struct {
	Channel string `json:"channel" query:"channel" validate:"required"`
}
