package config

type SiteInfo struct {
	CreateAt string `yaml:"create_at" json:"create_at"`
	Icp      string `yaml:"icp" json:"icp"`
	Title    string `yaml:"title" json:"title"`
	QQImage  string `yaml:"qq_image" json:"qq_image"`
	Version  string `yaml:"version" json:"version"`
	Email    string `yaml:"email" json:"email"`
	Wechat   string `yaml:"wechat" json:"wechat"`
	Name     string `yaml:"name" json:"name"`
	Job      string `yaml:"job" json:"job"`
	Addr     string `yaml:"addr" json:"addr"`
	Slogan   string `yaml:"slogan" json:"slogan"`
	SloganEn string `yaml:"slogan_en" json:"slogan_en"`
	Web      string `yaml:"web" json:"web"`
	Bilibili string `yaml:"bilibili" json:"bilibili"`
	Github   string `yaml:"github" json:"github"`
	Gitee    string `yaml:"gitee" json:"gitee"`
}
