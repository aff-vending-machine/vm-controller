package config

type RedisConfig struct {
	Host     string `default:"localhost" mapstructure:"HOST"`
	Port     int    `default:"6379" mapstructure:"PORT"`
	Username string `default:"" mapstructure:"USERNAME"`
	Password string `default:"" mapstructure:"PASSWORD"`
}
