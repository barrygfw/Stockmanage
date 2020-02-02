package common

import (
	"graduationProjectPeng/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Json_return(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	c.Abort()
}
