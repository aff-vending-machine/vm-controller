package config

type BootConfig struct {
	App       AppConfig       `mapstructure:"APP"`
	Fiber     FiberConfig     `mapstructure:"FIBER"`
	HTTP      HTTPConfig      `mapstructure:"HTTP"`
	RabbitMQ  RabbitMQConfig  `mapstructure:"RABBITMQ"`
	Redis     RedisConfig     `mapstructure:"REDIS"`
	SQLite    SQLiteConfig    `mapstructure:"SQLITE"`
	WebSocket WebSocketConfig `mapstructure:"WEBSOCKET"`
}
