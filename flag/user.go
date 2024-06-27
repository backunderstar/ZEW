package flag

import (
	"fmt"

	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/models"
	"github.com/backunderstar/zew/models/ctype"
	"github.com/backunderstar/zew/utils/pwd"
)

func CreateUser(permissions string) {
	var (
		username   string
		nickname   string
		password   string
		repassword string
		email      string
	)

	fmt.Printf("请输入用户名:")
	fmt.Scan(&username)
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickname)
	for {
		fmt.Printf("请输入密码:")
		fmt.Scan(&password)
		fmt.Printf("请再次输入密码:")
		fmt.Scan(&repassword)
		if password == repassword {
			break
		}
		fmt.Printf("两次密码不一致")
	}
	fmt.Printf("请输入邮箱:")
	fmt.Scan(&email)

	var userModel models.UserModel
	if global.DB.Take(&userModel, "email = ?", email).Error == nil {
		//	用户已存在
		fmt.Println("用户已存在")
		return
	}
	
	hashPwd := pwd.HashPwd(password)

	role := ctype.User
	if permissions == "admin" {
		role = ctype.Admin
	}

	err := global.DB.Create(&models.UserModel{
		Username:   username,
		Nickname:   nickname,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		SignStatus: ctype.Email,
		IP:        "127.0.0.1",
		Avatar:    "", 
	}).Error
	if err != nil {
		global.Log.Error("创建用户失败", err)
		return
	}
	global.Log.Infof("创建%s用户成功", username)
}
