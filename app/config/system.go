package config

type System struct {
	Mode    string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Host    string `mapstructure:"host" json:"host" yaml:"host"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	OssType string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	Domain  string `mapstructure:"domain" json:"domain" yaml:"domain"` // 访问domain
}
