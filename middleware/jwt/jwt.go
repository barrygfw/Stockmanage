package jwt

import (
	"github.com/gin-gonic/gin"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/util"
	"graduationProjectPeng/service/common"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Request.Header.Get("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			Claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > Claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			common.Json_return(c, code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}
