package request

type Get struct {
	ID uint `json:"id" query:"id" validate:"required"`
}
