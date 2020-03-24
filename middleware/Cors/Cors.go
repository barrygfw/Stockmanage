package Cors

import (
	"fmt"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/pkg/setting"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		origin := c.Request.Header.Get("Origin")
		filterHost := setting.ServerSetting.FilterHost
		var isAccess = false
		for _, v := range filterHost {
			if match, _ := regexp.MatchString(v, origin); match {
				isAccess = true
				break
			}
		}
		if isAccess {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			//c.Set("Content-Type", "application/json")
		} else {
			logging.Warn(origin, "域名未通过验证")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
