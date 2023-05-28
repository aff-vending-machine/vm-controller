package sqlite

import (
	"github.com/aff-vending-machine/vm-controller/configs"
	"github.com/aff-vending-machine/vm-controller/pkg/boot"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	*gorm.DB
	configs.SQLiteConfig
}

func New(cfg configs.SQLiteConfig) *Client {
	dsn := cfg.Database

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.LogLevel(cfg.LogLevel)),
	})
	boot.TerminateWhenError(err)

	return &Client{
		db,
		cfg,
	}
}
