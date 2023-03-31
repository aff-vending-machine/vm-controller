package config

type SmartEDCConfig struct {
	Port         string `default:"/dev/ttyACM0" mapstructure:"PORT"`
	TimeoutInSec int    `default:"300" mapstructure:"TIMEOUT_IN_SEC"`
}
