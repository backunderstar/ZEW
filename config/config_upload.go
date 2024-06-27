package config

type Upload struct {
	Size   float64    `yaml:"size" json:"size"`
	Path   string   `yaml:"path" json:"path"`
	Suffix []string `yaml:"suffix"` // 允许上传文件后缀
}
