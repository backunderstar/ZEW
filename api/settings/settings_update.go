package settings

import (
	"github.com/backunderstar/zew/config"
	"github.com/backunderstar/zew/core"
	"github.com/backunderstar/zew/global"
	"github.com/backunderstar/zew/utils/res"

	"github.com/gin-gonic/gin"
)

func (s *SettingsApi) SettingsUpdate(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParametersError, c)
		return
	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParametersError, c)
			return
		}
		global.Config.SiteInfo = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParametersError, c)
			return
		}
		global.Config.Jwt = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParametersError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParametersError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParametersError, c)
			return
		}
		global.Config.QiNiu = info
	default:
		res.FailWithMsg("没有对应的配置信息", c)
	}
	err = core.UpdateYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.JustOk(c)
}
