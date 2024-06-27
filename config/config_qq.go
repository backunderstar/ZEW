package config

import "fmt"

type QQ struct {
	AppID    string `yaml:"app_id" json:"app_id"`     // appid
	Key      string `yaml:"key" json:"key"`           // key
	Redirect string `yaml:"redirect" json:"redirect"` // 回调地址
}

func (q *QQ) GetPath() string {
	if q.AppID == "" || q.Key == "" || q.Redirect == "" {
		return ""
	}
	// https://q.qlogo.cn/g?b=qq&nk=111111&s=100
	return fmt.Sprintf("https://grapn.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_url=%s", q.AppID, q.Redirect)
}
