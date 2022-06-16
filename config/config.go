package config

type Config struct {
	Http  *Http  `mapstructure:"http" json:"http" yaml:"http"`    // HTTP服务相关配置
	Redis *Redis `mapstructure:"redis" json:"redis" yaml:"redis"` // Redis配置
	DB    *DB    `mapstructure:"db" json:"db" yaml:"db"`          // DB配置
	JWT   *JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
