package config

type RasPiConfig struct {
	LCDDevice            string `default:"/dev/fb0" mapstructure:"LCD_DEVICE"`
	LCDRotate            int    `default:"0" mapstructure:"LCD_ROTATE"`
	KeypadHorizontalLine []int  `default:"" mapstructure:"KEYPAD_HORIZONTAL_LINE"`
	KeypadVerticalLine   []int  `default:"" mapstructure:"KEYPAD_VERTICAL_LINE"`
}
