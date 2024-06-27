package main

import (
	"github.com/backunderstar/zew/core"
	"github.com/backunderstar/zew/flag"
	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/route"
)

func main() {
	// 初始化配置
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 初始化Redis
	/* global.Redis = core.InitRedis() */
	// 解析命令行参数
	option := flag.Parse()
	if flag.IsStopWeb(&option) {
		flag.SwitchOption(&option)
		return
	}
	// 初始化路由
	/* router := routers.InitRouter() */
	router := route.Routers()

	// 启动服务
	addr := global.Config.System.Addr()
	global.Log.Infof("服务将运行在: %s", addr)
	router.Run(addr)
}
