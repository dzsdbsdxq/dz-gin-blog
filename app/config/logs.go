package config

import "go.uber.org/zap/zapcore"

type Logs struct {
	Level      zapcore.Level `mapstructure:"level" json:"level" yaml:"level"`
	Path       string        `mapstructure:"path" json:"path" yaml:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize" yaml:"max-size"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups" yaml:"max-backups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`
	Compress   bool          `mapstructure:"compress" json:"compress" yaml:"compress"`
}
