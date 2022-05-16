package config

import "github.com/soonio/pupil/config/internal"

type Config struct {
	Http  *internal.Http  `mapstructure:"http" json:"http" yaml:"http"`    // HTTP服务相关配置
	Redis *internal.Redis `mapstructure:"redis" json:"redis" yaml:"redis"` // Redis配置
	DB    *internal.DB    `mapstructure:"db" json:"db" yaml:"db"`          // DB配置
}
