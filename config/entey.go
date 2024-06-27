package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Jwt      Jwt      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	QQ       QQ       `yaml:"qq"`
	QiNiu    QiNiu    `yaml:"qiniu"`
	Redis    Redis    `yaml:"redis"`
	Upload   Upload   `yaml:"upload"`
}
