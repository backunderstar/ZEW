package ctype

import "encoding/json"

type Role int

const (
	Admin Role = 1
	User  Role = 2
	Guest Role = 3
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	switch s {
	case Admin:
		return "管理员"
	case User:
		return "用户"
	case Guest:
		return "游客"
	default:
		return "Unknow"
	}
}
