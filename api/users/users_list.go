package users

import (
	"github.com/backunderstar/zew/models"
	"github.com/backunderstar/zew/models/ctype"
	"github.com/backunderstar/zew/service/common"
	"github.com/backunderstar/zew/utils/desens"
	"github.com/backunderstar/zew/utils/jwt"
	"github.com/backunderstar/zew/utils/res"
	"github.com/gin-gonic/gin"
)

func (u *UsersApi) UsersList(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ParametersError, c)
		return
	}
	list, count, _ := common.ComSingleList(models.UserModel{}, common.Option{
		PageInfo: page,
	})

	var users []models.UserModel
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.Admin {
			user.Username = ""
			user.Email = desens.DesensitizationEmail(user.Email)
		}
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}
