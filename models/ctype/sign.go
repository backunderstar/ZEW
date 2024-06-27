package ctype

import "encoding/json"

// SignStatus 定义了登录状态的枚举类型。
// 它用于表示用户登录的方式，比如通过QQ或电子邮件。
type SignStatus int

// 枚举常量定义了具体的登录状态。
const (
	// QQ 表示通过QQ方式登录。
	QQ SignStatus = 0
	// Email 表示通过电子邮件方式登录。
	Email SignStatus = 1
)

// MarshalJSON 实现了 json.Marshaler 接口。
// 它将 SignStatus 类型的值转换为对应的字符串表示，用于 JSON 序列化。
func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// String 实现了 fmt.Stringer 接口。
// 它将 SignStatus 类型的值转换为人类可读的字符串表示。
func (s SignStatus) String() string {
	switch s {
	case QQ:
		return "QQ"
	case Email:
		return "Email"
	default:
		return "Unknown"
	}
}
