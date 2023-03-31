package repository

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func MakeQuery(db *gorm.DB, filter []string) *gorm.DB {
	tx := db

	for _, filter := range filter {
		f := strings.Split(filter, ":")

		if len(f) != 3 {
			// invalid filter
			continue
		}

		field := f[0]
		condition := f[1]
		value := f[2]

		switch strings.ToUpper(condition) {
		case "LIMIT":
			if limit, err := strconv.Atoi(value); err == nil {
				tx = tx.Limit(limit)
			}

		case "OFFSET":
			if offset, err := strconv.Atoi(value); err == nil {
				tx = tx.Offset(offset)
			}

		case "ORDER":
			if value == "asc" {
				tx = tx.Order(field)
			} else {
				order := fmt.Sprintf("%s %s", field, value)
				tx = tx.Order(order)
			}

		case "PRELOAD":
			tx = tx.Preload(value)

		case "BETWEEN":
			where := fmt.Sprintf("%s BETWEEN ? AND ?", field)
			tx = tx.Where(where, f[2], f[3])

		default:
			where := fmt.Sprintf("%s %s ?", field, condition)
			tx = tx.Where(where, value)
		}
	}

	return tx
}
