package configs

type HTTPConfig struct {
	Host         string `default:"localhost" mapstructure:"HOST"`
	Cert         bool   `default:"" mapstructure:"CERT"`
	TimeoutInSec int    `default:"30" mapstructure:"TIMEOUT_IN_SEC"`
}
