package repository

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func MakeQuery(db *gorm.DB, filter []string) *gorm.DB {
	tx := db

	for _, filter := range filter {
		f := strings.Split(filter, ":")

		if len(f) != 3 && len(f) != 4 {
			log.Warn().Int("length", len(f)).Str("filter", filter).Msg("unable to make query")
			continue
		}

		field := f[0]
		condition := f[1]
		value := f[2]
		type_ := "string"

		if len(f) == 4 {
			type_ = f[3]
		}

		switch strings.ToUpper(condition) {
		case "OFFSET":
			if offset, ok := ToValue(value, "int"); ok {
				tx = tx.Offset(offset.(int))
			} else {
				log.Warn().Str("filter", filter).Msg("unable to OFFSET")
			}

		case "LIMIT":
			if limit, ok := ToValue(value, "int"); ok {
				tx = tx.Limit(limit.(int))
			} else {
				log.Warn().Str("filter", filter).Msg("unable to LIMIT")
			}

		case "SORT":
			if value == "asc" {
				tx = tx.Order(field)
			} else {
				order := fmt.Sprintf("%s %s", field, value)
				tx = tx.Order(order)
			}

		case "SELECT":
			queries := strings.Split(value, ",")
			tx.Select(queries)

		case "PRELOAD":
			tx = tx.Preload(value)

		case "BETWEEN":
			where := fmt.Sprintf("%s BETWEEN ? AND ?", field)

			val := strings.Split(value, ",")

			if len(val) != 2 {
				log.Warn().Int("length", len(val)).Str("filter", filter).Msg("unable to BETEWEEN")
				continue
			}

			v1, ok1 := ToValue(val[0], type_)
			v2, ok2 := ToValue(val[1], type_)

			if ok1 && ok2 {
				tx = tx.Where(where, v1, v2)
			} else {
				log.Warn().Bool("value 1", ok1).Bool("value 2", ok2).Str("filter", filter).Msg("unable to BETEWEEN")
			}

		default:
			where := fmt.Sprintf("%s %s ?", field, condition)
			if val, ok := ToValue(value, type_); ok {
				tx = tx.Where(where, val)
			} else {
				log.Warn().Str("filter", filter).Msg("unable to WHERE")
			}
		}
	}

	return tx
}

func ToValue(v string, t string) (interface{}, bool) {
	switch t {
	case "string":
		return v, true

	case "int":
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Warn().Err(err).Str("value", v).Msg("unable to parse int")
			return nil, false
		}
		return n, true

	case "uint":
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Warn().Err(err).Str("value", v).Msg("unable to parse uint")
			return nil, false
		}
		return n, true

	case "[]uint":
		n := make([]uint, 0)
		err := json.Unmarshal([]byte(v), &n)
		if err != nil {
			log.Warn().Err(err).Str("value", v).Msg("unable to parse []uint")
			return nil, false
		}
		return n, true

	case "bool":
		return strings.EqualFold(v, "true"), true

	case "time":
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, v)
		if err != nil {
			log.Warn().Err(err).Str("value", v).Msg("unable to parse time")
			return nil, false
		}

		return t, true
	}

	return nil, false
}
