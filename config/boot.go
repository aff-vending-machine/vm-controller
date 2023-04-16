package config

type BootConfig struct {
	App       AppConfig       `mapstructure:"APP"`
	Fiber     FiberConfig     `mapstructure:"FIBER"`
	HTTP      HTTPConfig      `mapstructure:"HTTP"`
	Board     BoardConfig     `mapstructure:"BOARD"`
	RabbitMQ  RabbitMQConfig  `mapstructure:"RABBITMQ"`
	Redis     RedisConfig     `mapstructure:"REDIS"`
	SQLite    SQLiteConfig    `mapstructure:"SQLITE"`
	WebSocket WebSocketConfig `mapstructure:"WEBSOCKET"`
}
