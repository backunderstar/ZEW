package settings

import (
	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/utils/res"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (s *SettingsApi) SettingsInfo(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParametersError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMsg("没有对应的配置信息", c)
	}
}
