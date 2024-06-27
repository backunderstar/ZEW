package common

import (
	"fmt"

	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/models"
	"gorm.io/gorm"
)

// Option 扩展了分页信息，加入了模糊搜索和自定义条件
type Option struct {
	models.PageInfo
	SearchFields map[string]interface{}     // 模糊搜索字段
	Preload      []string                   // 需要预加载的关联关系
	Where        func(tx *gorm.DB) *gorm.DB // 自定义WHERE条件
}

// ComSingleList 公用获取列表数据方法（适用于单表），支持模糊匹配、预加载和自定义WHERE条件
func ComSingleList[T any](model T, option Option) (list []T, count int64, err error) {
	// 设置默认值
	if option.Page == 0 {
		option.Page = 1
	}
	if option.Limit == 0 {
		option.Limit = 10
	}
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	// 构建查询对象
	db := global.DB.Model(&model)

	// 应用自定义WHERE条件
	if option.Where != nil {
		db = option.Where(db)
	}

	// 模糊搜索
	for key, value := range option.SearchFields {
		db = db.Where(key+" LIKE ?", "%"+fmt.Sprint(value)+"%")
	}

	// 预加载关联关系
	for _, preloadField := range option.Preload {
		db = db.Preload(preloadField)
	}

	// 执行带分页的查询
	err = db.Limit(option.Limit).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	// 分离计数逻辑，正确获取总数
	var totalCount int64
	err = db.Count(&totalCount).Error
	if err != nil {
		return list, 0, err
	}
	count = totalCount
	return list, count, err
}
