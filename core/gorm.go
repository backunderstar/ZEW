package core

import (
	"time"
	"github.com/backunderstar/zew/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置MySQL,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.DSN()

	var mySqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mySqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mySqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mySqlLogger,
	})
	if err != nil {
		global.Log.Fatalf("数据库连接失败: %v", err)
	} else {
		global.Log.Infoln("数据库连接成功")
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 2)
	return db
}
