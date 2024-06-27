package models

import (
	"github.com/backunderstar/zew/models/ctype"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title         string         `gorm:"size:64" json:"title"`                    // 文章标题
	Abstract      string         `json:"abstract"`                                // 文章摘要
	Content       string         `gorm:"type:longtext" json:"content"`            // 文章内容
	LookCount     int            `gorm:"default:0" json:"look_count"`             // 浏览量
	CommentCount  int            `gorm:"default:0" json:"comment_count"`          // 评论量
	DiggCount     int            `gorm:"default:0" json:"digg_count"`             // 点赞量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tag_models"` // 文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`           // 文章评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`              // 文章作者
	UserID        uint           `json:"user_id"`                                 // 用户
	Category      string         `gorm:"size:20" json:"category"`                 // 文章分类
	Source        string         `json:"source"`                                  // 文章来源
	Link          string         `json:"link"`                                    // 文章来源链接
	Cover         ImageModel     `json:"-"`                                       // 文章封面
	CoverID       uint           `json:"cover_id"`                                // 文章封面ID
	CoverPath     string         `json:"cover_path"`                              // 文章封面路径
	Nickname      string         `gorm:"size:42" json:"nick_name"`                // 文章作者昵称
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`         // 文章标签
}
