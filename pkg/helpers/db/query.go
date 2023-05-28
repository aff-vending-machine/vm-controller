package db

import (
	"reflect"
	"strings"
)

type Query struct {
	Limit    *int
	Offset   *int
	Where    []Where
	Order    *Order
	Perloads []Preload // gorm only
}

type Where struct {
	Query string
	Args  []interface{}
}

type Order struct {
	Field     string
	Decending bool
}

type Preload struct {
	Query string
	Args  []interface{}
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) PtrWhere(query string, args ...interface{}) *Query {
	// Return early if there are no arguments
	if len(args) == 0 {
		return q
	}

	// Use a type switch to handle different types
	value := reflect.ValueOf(args[0])
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return q
	}

	// Add the where clause
	q.Where = append(q.Where, Where{Query: query, Args: args})
	return q
}

func (q *Query) PtrOrder(order *string) *Query {
	if order == nil {
		return q
	}

	sortBy := strings.Split(*order, ":")
	decending := len(sortBy) > 1 && strings.ToLower(sortBy[1]) == "desc"

	q.Order = &Order{Field: sortBy[0], Decending: decending}
	return q
}

func (q *Query) PtrLimit(limit *int) *Query {
	q.Limit = limit
	return q
}

func (q *Query) PtrOffset(offset *int) *Query {
	q.Offset = offset
	return q
}

func (q *Query) PtrPreloads(preloads *string) *Query {
	// Return early if there are no arguments
	if preloads == nil {
		return q
	}

	list := strings.Split(*preloads, ":")
	for _, preload := range list {
		q.AddPreload(preload)
	}

	// Add the where clause
	return q
}

func (q *Query) AddWhere(query string, args ...interface{}) *Query {
	q.Where = append(q.Where, Where{Query: query, Args: args})
	return q
}

func (q *Query) AddPreload(query string, args ...interface{}) *Query {
	q.Perloads = append(q.Perloads, Preload{Query: query, Args: args})
	return q
}

func (q *Query) SetLimit(limit int) *Query {
	q.Limit = &limit
	return q
}

func (q *Query) SetOffset(offset int) *Query {
	q.Offset = &offset
	return q
}

func (q *Query) SetOrder(order string) *Query {
	sortBy := strings.Split(order, ":")
	decending := len(sortBy) > 1 && strings.ToLower(sortBy[1]) == "desc"

	q.Order = &Order{Field: sortBy[0], Decending: decending}
	return q
}

func (q *Query) Clear() *Query {
	q.Where = []Where{}
	q.Limit = nil
	q.Offset = nil
	q.Perloads = []Preload{}
	return q
}
