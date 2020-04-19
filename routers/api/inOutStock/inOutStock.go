package inOutStock

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/inOutStockModel"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/service/common"
	"graduationProjectPeng/service/inOutStockService"

	"github.com/gin-gonic/gin"
)

/**
出入库
api : /api/stock/inoutstock
params : json
{
	"type": 1,
	"goods_id", 1,
	"num", 1000,
	"comment": "从xx厂进货1000件"
}
*/
func GoodsInOutStock(c *gin.Context) {
	var inOutStock inOutStockModel.InoutStock
	if err := c.ShouldBindJSON(&inOutStock); err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if inOutStock.Type != inOutStockModel.OutputStock && inOutStock.Type != inOutStockModel.InputStock {
		logging.Warn("出入库类型参数错误")
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if inOutStock.Num < 1 {
		logging.Warn("出入库数量参数错误")
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if isSucs, code := inOutStockService.AddInOutStockRow(&inOutStock); !isSucs {
		logging.Warn(e.GetMsg(code))
		common.Json_return(c, code, "")
		return
	}
	logging.Info(inOutStock)
	common.Json_return(c, e.SUCCESS, "")
}

func QueryInoutList(c *gin.Context) {
	var param models.InoutListParam
	if err := c.ShouldBindQuery(&param); err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if param.Type != 0 && param.Type != inOutStockModel.InputStock && param.Type != inOutStockModel.OutputStock {
		logging.Warn("出入库类型参数错误")
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	data, err := inOutStockService.QueryInoutStockList(&param)
	if err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.ERROR, "")
		return
	}
	common.Json_return(c, e.SUCCESS, data)
}
