package request

type Filter struct {
	Limit  *int    `json:"limit,omitempty" query:"limit"`
	Offset *int    `json:"offset,omitempty" query:"offset"`
	SortBy *string `json:"sort_by,omitempty" query:"sort_by"`
	Search *string `json:"search,omitempty" query:"search"`
}
