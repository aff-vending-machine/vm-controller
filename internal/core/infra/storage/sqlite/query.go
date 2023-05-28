package sqlite

import (
	"github.com/aff-vending-machine/vm-controller/pkg/db"

	"gorm.io/gorm"
)

func MakeQuery(db *gorm.DB, query *db.Query) *gorm.DB {
	if query == nil {
		return db
	}

	if query.Limit != nil && *query.Limit > 0 {
		db = db.Limit(*query.Limit)
	}

	if query.Offset != nil && *query.Offset > 0 {
		db = db.Offset(*query.Offset)
	}

	if len(query.Where) > 0 {
		for _, where := range query.Where {
			db = db.Where(where.Query, where.Args...)
		}
	}

	if query.Order != nil {
		if query.Order.Decending {
			db = db.Order(query.Order.Field + " DESC")
		} else {
			db = db.Order(query.Order.Field)
		}
	}

	if len(query.Perloads) > 0 {
		for _, perload := range query.Perloads {
			db = db.Preload(perload.Query, perload.Args...)
		}
	}

	return db
}
