package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	Title    string         `gorm:"size:32" json:"title"`           // 标签名
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"-"` // 标签对应的文章
}
