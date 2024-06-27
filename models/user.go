package models

import (
	"github.com/backunderstar/zew/models/ctype"
	"gorm.io/gorm"
)

// UserModel 用户表
type UserModel struct {
	gorm.Model
	Nickname      string           `gorm:"size:36" json:"nickname"`             // 昵称
	Username      string           `gorm:"size:36" json:"username"`             // 用户名
	Password      string           `gorm:"size:128" json:"-"`                   // 密码
	Avatar        string           `gorm:"size:256" json:"avatar"`              // 头像
	Email         string           `gorm:"size:128" json:"email"`               // 手机号
	Addr          string           `gorm:"size:64" json:"address"`              // 地址
	Token         string           `gorm:"size:64" json:"token"`                // token 后续写到redis中，不写入这里了，原来是想做永久登录的
	IP            string           `gorm:"size:20" json:"ip"`                   // ip
	Role          ctype.Role       `gorm:"size:4;default:1" json:"role"`        // 角色 1 管理员 2 普通用户 3 游客 4 被禁用
	SignStatus    ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"` // 注册来源 qq 邮箱 手机等
	ArticleModels []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`          // 发布文章列表
}
