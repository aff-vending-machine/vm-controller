package configs

type SQLiteConfig struct {
	LogLevel int    `default:"1" mapstructure:"LOG_LEVEL"`
	Database string `default:"db/default.sqlite" mapstructure:"DATABASE"`
}
