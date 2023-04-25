package request

import (
	"encoding/json"
	"fmt"
)

type Clear struct {
	Query struct {
		IDs []uint
	}
}

func (r *Clear) ToFilter() []string {
	s, _ := json.Marshal(r.Query.IDs)

	return []string{
		fmt.Sprintf("id:IN:%s:[]uint", s),
	}
}
