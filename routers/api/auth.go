package api

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/pkg/setting"
	"graduationProjectPeng/pkg/util"
	"graduationProjectPeng/service/common"

	"github.com/gin-gonic/gin"
)

/**
用户登录
api : api/login
params : json
{
	"username":"xxxxx",
	"password":"xxxxx"
}
*/
func Login(c *gin.Context) {
	var userinfo models.User
	var code int
	data := make(map[string]string)
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		code = e.INVALID_PARAMS
		logging.Info(e.GetMsg(code))
		common.Json_return(c, code, data)
		return
	}
	if userinfo.Username != setting.UserSetting.Name {
		code = e.ERROR_USER_NONE
	} else if userinfo.Password != setting.UserSetting.Pass {
		code = e.ERROR_USER_PASS
	} else {
		token, err := util.GenerateToken(userinfo.Username)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	}
	logging.Info(e.GetMsg(code))
	common.Json_return(c, code, data)
}
