package config

type PostgreSQLConfig struct {
	Enable   bool   `default:"false" mapstructure:"ENABLE"`
	LogLevel int    `default:"1" mapstructure:"LOG_LEVEL"`
	Username string `default:"" mapstructure:"USERNAME"`
	Password string `default:"" mapstructure:"PASSWORD"`
	Host     string `default:"localhost" mapstructure:"HOST"`
	Port     int    `default:"5432" mapstructure:"PORT"`
	Database string `default:"postgresdb" mapstructure:"DATABASE"`
	SSLMode  string `default:"disable" mapstructure:"SSLMODE"`
	TimeZone string `default:"Asia/Bangkok" mapstructure:"TIMEZONE"`
}
