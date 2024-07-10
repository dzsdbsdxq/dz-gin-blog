package config

type Local struct {
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // 本地文件存储路径
	BaseUrl   string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`       // 本地文件访问URL
}
