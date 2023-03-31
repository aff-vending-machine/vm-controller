package repository

import (
	"gorm.io/gorm"
)

type Template[T any] struct {
	DB *gorm.DB
}

func New[T any](db *gorm.DB) *Template[T] {
	return &Template[T]{DB: db}
}
