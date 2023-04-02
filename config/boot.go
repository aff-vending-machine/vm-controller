package config

type BootConfig struct {
	App    AppConfig    `mapstructure:"APP"`
	Fiber  FiberConfig  `mapstructure:"FIBER"`
	HTTP   HTTPConfig   `mapstructure:"HTTP"`
	RasPi  RasPiConfig  `mapstructure:"RASPI"`
	Redis  RedisConfig  `mapstructure:"REDIS"`
	SQLite SQLiteConfig `mapstructure:"SQLITE"`
}
