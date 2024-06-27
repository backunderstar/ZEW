package ctype

import (
	"database/sql/driver"
	"strings"
)

// Array 是一个字符串数组类型，用于在数据库操作中存储和检索字符串列表。
type Array []string

// Scan 是Array类型的Scanner接口实现。
// 它负责从数据库的列值中解析并存储字符串数组。
// 参数value是数据库查询结果中列的值，它被期望是一个[]byte类型。
// 函数将byte切片转换为字符串，分割成多行，并存储到Array中。
func (t *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(string(v), "\n")
	return nil
}

// Value 是Array类型的driver.Valuer接口实现。
// 它负责将Array转换为数据库可以接受的值。
// 返回值是一个字符串，多个字符串用换行符("\n")连接。
func (t Array) Value() (driver.Value, error) {
	return strings.Join(t, "\n"), nil
}