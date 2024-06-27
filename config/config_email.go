package config

type Email struct {
	Host             string `yaml:"host" json:"host"`
	Port             int    `yaml:"port" json:"port"`
	User             string `yaml:"user" json:"user"`                             // 发件人邮箱
	Password         string `yaml:"password" json:"password"`                     // 发件人邮箱密码
	DefaultFromEmail string `yaml:"default_from_email" json:"default_from_email"` // 默认发件人
	UserSSL          bool   `yaml:"user_ssl" json:"user_ssl"`                     // 是否启用SSL
	UserTLS          bool   `yaml:"user_tls" json:"user_tls"`                     // 是否启用TLS
}
