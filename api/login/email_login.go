package login

import (
	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/models"
	"github.com/backunderstar/zew/utils/jwt"
	"github.com/backunderstar/zew/utils/pwd"
	"github.com/backunderstar/zew/utils/res"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	Username string `json:"username" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (l *LoginApi) EmailLogin(c *gin.Context) {
	var req EmailLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithCode(1002, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "username = ? OR email = ?", req.Username, req.Username).Error
	if err != nil {
		global.Log.Warn("用户名不存在", err)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	isCheck := pwd.Verify(user.Password, req.Password)
	if !isCheck {
		global.Log.Warn("密码错误", err)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwt.GenerateToken(jwt.JwtPayLoad{
		Nickname: user.Nickname,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error("生成token失败", err)
		res.FailWithMsg("生成token失败", c)
		return
	}
	res.OkWithData(token, c)

}
