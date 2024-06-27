package flag

import (
	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/models"
)

// MakeMigration 实现数据库的自动迁移
// 该函数通过扫描指定的模型结构，自动创建对应的数据表，用于初始化数据库结构或更新数据库结构到最新版本。
func MakeMigration() {
	var err error

	// 设置关联表的结构，这里演示了如何设置菜单模型与菜单图片模型的关联关系
	/* global.DB.SetupJoinTable(&models.UserModel{}, "collectsModels", &models.UserCollectsModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Image", &models.MenuImageModel{}) */

	// 开始日志记录，表示数据库迁移过程开始
	global.Log.Infof("开始迁移数据库")
	// 执行自动迁移，指定使用InnoDB引擎，并列出所有需要迁移的模型
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.ArticleModel{},
		&models.TagModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.MenuModel{},
		&models.LoginDataModel{},
		&models.ImageModel{},
		&models.MessageModel{},
	)
	// 检查迁移过程中是否有错误发生，如果有，则记录错误日志并返回
	if err != nil {
		global.Log.Errorf("数据库迁移失败: %v", err)
		return
	}
	// 记录成功日志，表示数据库迁移完成
	global.Log.Info("数据库迁移成功")
}
