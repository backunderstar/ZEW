package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` // 连接参数 Charset
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级:debug、release
}

func (m *Mysql) DSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + fmt.Sprint(m.Port) + ")/" + m.Db + "?" + m.Config
}
