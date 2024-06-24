package config

type System struct {
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
