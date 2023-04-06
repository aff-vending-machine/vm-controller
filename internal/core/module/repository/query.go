package repository

import (
	"encoding/json"
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

			if val, ok := strings.CutSuffix(value, ".([]uint)"); ok {
				uints := make([]uint, 0)
				json.Unmarshal([]byte(val), &uints)
				tx = tx.Where(where, uints)
				continue
			}

			if val, ok := strings.CutSuffix(value, ".(bool)"); ok {
				tx = tx.Where(where, strings.HasPrefix(val, "true"))
				continue
			}

			if val, ok := strings.CutSuffix(value, ".(int)"); ok {
				if num, err := strconv.Atoi(val); err != nil {
					tx = tx.Where(where, num)
					continue
				}
			}

			tx = tx.Where(where, value)
		}
	}

	return tx
}
