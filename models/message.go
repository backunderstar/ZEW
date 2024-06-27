package models

import "gorm.io/gorm"

// MessageModel 记录消息
type MessageModel struct {
	gorm.Model
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"` // 发送者id
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickname string    `gorm:"size:42" json:"send_user_nickname"`
	SendUserAvatar   string    `json:"send_user_avatar"`

	RevUserID       uint      `gorm:"primaryKey" json:"rev_user_id"` // 接收者id
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickname string    `gorm:"size:42" json:"rev_user_nickname"`
	RevUserAvatar   string    `json:"rev_user_avatar"`

	IsRead  bool   `gorm:"default:false" json:"is_read"` // 接收方是否已读
	Content string `gorm:"type:longtext" json:"content"` // 消息内容
}
