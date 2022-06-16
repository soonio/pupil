package config

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"` // jwt签名
	Duration   int64  `mapstructure:"duration" json:"duration" yaml:"duration"`         // 生命周期
	Issuer     string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`               // 签发者
}
