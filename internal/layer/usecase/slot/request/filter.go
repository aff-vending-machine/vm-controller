package request

type Filter struct {
	Limit  *int `json:"limit,omitempty" query:"limit"`
	Offset *int `json:"offset,omitempty" query:"offset"`
}
