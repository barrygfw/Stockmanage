package routers

import (
	"graduationProjectPeng/middleware/jwt"
	"graduationProjectPeng/pkg/setting"
	"graduationProjectPeng/routers/api"
	"graduationProjectPeng/routers/api/category"
	"graduationProjectPeng/routers/api/goods"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.POST("/api/login", api.Login)
	root := r.Group("/api")
	if setting.AppSetting.CheckToken {
		root.Use(jwt.JWT())
	}
	{
		categoryApi := root.Group("/category")
		{
			categoryApi.GET("/getall", category.GetCate)
			categoryApi.POST("/add", category.AddCate)
			categoryApi.POST("/del", category.DelCate)
			categoryApi.POST("/update", category.UpdateCate)
		}
		goodsApi := root.Group("/goods")
		{
			goodsApi.POST("/add", goods.AddGoods)
			goodsApi.POST("/update", goods.UpdateGoods)
			goodsApi.POST("/del", goods.DelGoods)
		}
	}
}