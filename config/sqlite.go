package config

type SQLiteConfig struct {
	Enable   bool   `default:"false" mapstructure:"ENABLE"`
	LogLevel int    `default:"1" mapstructure:"LOG_LEVEL"`
	Database string `default:"db/default.sqlite" mapstructure:"DATABASE"`
}
