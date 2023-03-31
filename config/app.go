package config

type AppConfig struct {
	ENV      string        `default:"production" mapstructure:"ENV"`
	LogLevel int           `default:"0" mapstructure:"LOG_LEVEL"`
	Preload  bool          `default:"false" mapstructure:"PRELOAD"`
	Jaeger   bool          `default:"false" mapstructure:"JAEGER"`
	Asset    AssetConfig   `mapstructure:"ASSET"`
	Center   CenterConfig  `mapstructure:"CENTER"`
	Machine  MachineConfig `mapstructure:"MACHINE"`
}

type AssetConfig struct {
	ImagePath string `default:"./assets/images/1440x2560" mapstructure:"IMAGE_PATH"`
	FontPath  string `default:"./assets/fonts" mapstructure:"FONT_PATH"`
}

type CenterConfig struct {
	RPCQueue         string `default:"CT-APP-RPC" mapstructure:"RPC_QUEUE"`
	TopicQueue       string `default:"CT-APP-TOPIC" mapstructure:"TOPIC_QUEUE"`
	RoutingKeyHeader string `default:"routing-key" mapstructure:"ROUTING_KEY_HEADER"`
}

type MachineConfig struct {
	Name         string `default:"" mapstructure:"NAME"`
	SerialNumber string `default:"" mapstructure:"SERIAL_NUMBER"`
	Location     string `default:"unknown" mapstructure:"LOCATION"`
	Type         string `default:"unknown" mapstructure:"TYPE"`
	Vendor       string `default:"unknown" mapstructure:"VENDOR"`
	Center       string `default:"CT-APP-CENTER" mapstructure:"CENTER"`
}
