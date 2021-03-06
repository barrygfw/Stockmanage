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
	"graduationProjectPeng/routers/api/toDoList"
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
			goodsApi.GET("/query", goods.QueryGoods)
		}
		inOutStockApi := root.Group("/stock")
		{
			inOutStockApi.POST("/inoutstock", inOutStock.GoodsInOutStock)
			inOutStockApi.GET("/inoutlist", inOutStock.QueryInoutList)
			inOutStockApi.POST("/inventory", inOutStock.StockInventory)
			inOutStockApi.GET("/inventorylist", inOutStock.QueryInventory)
		}
		uploadApi := root.Group("/upload")
		{
			uploadApi.POST("/image", api.UploadImage)
		}
		toDoListApi := root.Group("/todo")
		{
			toDoListApi.POST("/add", toDoList.AddTodo)
			toDoListApi.POST("/update", toDoList.UptTodo)
			toDoListApi.POST("/del", toDoList.DelTodo)
			toDoListApi.GET("/query", toDoList.Query)
		}
	}
}
