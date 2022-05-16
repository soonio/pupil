package config

type Config struct {
	Http *Http `mapstructure:"http" json:"http" yaml:"http"` // HTTP服务相关配置
}
