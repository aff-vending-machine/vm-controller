package request

type GetOne struct {
	Channel string `json:"channel" query:"channel" validate:"required"`
}
