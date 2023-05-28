package configs

type RabbitMQConfig struct {
	Protocol string `default:"amqp" mapstructure:"PROTOCOL"`
	Host     string `default:"localhost" mapstructure:"HOST"`
	Port     string `default:"5672" mapstructure:"PORT"`
	Username string `default:"" mapstructure:"USERNAME"`
	Password string `default:"" mapstructure:"PASSWORD"`
	Path     string `default:"" mapstructure:"PATH"`
}
