package config

type Config struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Logs    Logs    `mapstructure:"logs" json:"logs" yaml:"logs"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors    CORS    `mapstructure:"cors" json:"cors" yaml:"cors"`
	Local   Local   `mapstructure:"local" json:"local" yaml:"local"`
}
