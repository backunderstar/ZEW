package api

import (
	"github.com/backunderstar/zew/api/login"
	"github.com/backunderstar/zew/api/settings"
	"github.com/backunderstar/zew/api/upload"
	"github.com/backunderstar/zew/api/users"
)

type ApiGroup struct {
	LoginApi    login.LoginApi
	SettingsApi settings.SettingsApi
	UploadApi   upload.UploadApi
	UsersApi    users.UsersApi
}

var ApiGroupApp = new(ApiGroup)
