package sqlite

import (
	"github.com/aff-vending-machine/vmc-rpi-ctrl/config"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/boot"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg config.SQLiteConfig) *gorm.DB {
	dsn := cfg.Database

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.LogLevel(cfg.LogLevel)),
	})
	boot.TerminateWhenError(err)

	return db
}
