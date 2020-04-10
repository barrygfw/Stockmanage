package routers

import (
	"graduationProjectPeng/middleware/Cors"
	"graduationProjectPeng/middleware/jwt"
	"graduationProjectPeng/pkg/setting"
	"graduationProjectPeng/pkg/upload"
	"graduationProjectPeng/routers/api"
	"graduationProjectPeng/routers/api/category"
	"graduationProjectPeng/routers/api/goods"
	"graduationProjectPeng/routers/api/inOutStock"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.Use(Cors.Cors())
	r.StaticFS("/cdn/images", http.Dir(upload.GetImagePath()))
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
		inOutStockApi := root.Group("/stock")
		{
			inOutStockApi.POST("/inoutstock", inOutStock.GoodsInOutStock)
		}
		uploadApi := root.Group("/upload")
		{
			uploadApi.POST("/image", api.UploadImage)
		}
	}
}
