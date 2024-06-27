package config

type QiNiu struct {
	Enable    bool    `yaml:"enable" json:"enable"`
	AccessKey string  `yaml:"access_key" json:"access_key"` // AccessKey
	SecretKey string  `yaml:"secret_key" json:"secret_key"` // SecretKey
	Bucket    string  `yaml:"bucket" json:"bucket"`         // 存储空间名称
	CDN       string  `yaml:"cdn" json:"cdn"`               // 存储空间绑定的 CDN 加速域名
	Zone      string  `yaml:"zone" json:"zone"`             // 存储地区
	Size      float64 `yaml:"size" json:"size"`             // 上传文件大小限制，单位MB
}
