package flag

import (
	sys_flag "flag"

	"github.com/backunderstar/zew/global"
)

// Option 结构体用于存储程序的配置选项
type Option struct {
	DB   bool
	User string // -u admin ; -u user
}

// Parse 函数用于解析命令行参数，并返回程序的配置选项
// 参数: 无
// 返回值: Option - 程序配置选项
// 解析命令行： go run main.go -db
func Parse() Option {
	// 定义名为"db"的命令行参数，默认值为false，描述为"初始化数据库"
	// go run main.go -db
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	// 解析命令行参数
	sys_flag.Parse()
	// 返回配置选项
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsStopWeb 函数用于判断是否需要停止Web项目
// 参数: option *Option - 程序配置选项
// 返回值: bool - 是否需要停止Web项目
// 如果配置选项中的DB为true，则记录日志并返回true，表示需要停止Web项目；否则返回false
func IsStopWeb(option *Option) bool {
	if option.DB {
		global.Log.Infof("停止web项目")
		return true
	}
	if option.User != "" {
		global.Log.Infof("停止web项目")
		return true
	}
	return false
}

// SwitchOption 函数根据配置选项执行相应的操作
// 参数: option *Option - 程序配置选项
// 如果配置选项中的DB为true，则执行数据库迁移操作；否则显示使用帮助信息
func SwitchOption(option *Option) {
	if option.DB {
		// 迁移数据库
		MakeMigration()
		return
	}

	//fmt.Println("创建用户", option.User, option.User == "")
	if option.User == "admin" || option.User == "user" {
		// 创建用户
		CreateUser(option.User)
		return
	}
	sys_flag.Usage()
}
