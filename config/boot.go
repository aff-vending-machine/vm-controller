package config

type BootConfig struct {
	App        AppConfig        `mapstructure:"APP"`
	Fiber      FiberConfig      `mapstructure:"FIBER"`
	PostgreSQL PostgreSQLConfig `mapstructure:"POSTGRESQL"`
	Redis      RedisConfig      `mapstructure:"REDIS"`
	Mail       MailConfig       `mapstructure:"MAIL"`
	SmartEDC   SmartEDCConfig   `mapstructure:"SMARTEDC"`
	SQLite     SQLiteConfig     `mapstructure:"SQLITE"`
	WebSocket  WebSocketConfig  `mapstructure:"WEBSOCKET"`
}
