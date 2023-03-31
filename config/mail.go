package config

type MailConfig struct {
	Host     string `default:"" mapstructure:"HOST"`
	Port     int    `default:"" mapstructure:"PORT"`
	Username string `default:"" mapstructure:"USERNAME"`
	Password string `default:"" mapstructure:"PASSWORD"`
}
