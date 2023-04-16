package config

type WebSocketConfig struct {
	Port string `default:"8080" mapstructure:"PORT"`
}
