package goods

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/goodsModel"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/service/common"
	"graduationProjectPeng/service/goodsService"

	"github.com/gin-gonic/gin"
)

/**
新增商品
api : /api/goods/add
params : json
{
	"category_id":1, //所属分类id
	"name":"商品名称"
}
*/
func AddGoods(c *gin.Context) {
	var goods goodsModel.Goods
	//参数校验
	if err := c.ShouldBindJSON(&goods); err != nil {
		logging.Info(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, err.Error())
		return
	}
	if err := goodsService.InsertGoods(&goods); err != nil {
		logging.Fatal(err.Error())
		common.Json_return(c, e.ERROR, err.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

/**
更新商品
api : /api/goods/update
params : json
{
	"goods_id":1, //商品id
	"name":"商品名称",
	"category_id": 10 //分类id
}
*/
func UpdateGoods(c *gin.Context) {
	var goods goodsModel.Goods
	//参数校验
	if err := c.ShouldBindJSON(&goods); err != nil || goods.GoodsId == 0 {
		logging.Info(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, err.Error())
		return
	}
	if err := goodsService.UpdateGoods(&goods); err != nil {
		logging.Error(err.Error())
		common.Json_return(c, e.ERROR, err.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

/**
删除商品（支持批量）
api : /api/goods/del
params : json
{
	"ids" : [1,2,3] //商品id数组
}
*/
func DelGoods(c *gin.Context) {
	var goodsIds models.IdList
	//参数校验
	if err := c.ShouldBindJSON(&goodsIds); err != nil {
		logging.Info(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, err.Error())
		return
	}
	if err := goodsService.DeleteGoods(goodsIds.Ids); err != nil {
		logging.Error(err.Error())
		common.Json_return(c, e.ERROR, err.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}