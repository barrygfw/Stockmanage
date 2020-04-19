package inOutStock

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/InventoryModel"
	"graduationProjectPeng/models/inOutStockModel"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/service/common"
	"graduationProjectPeng/service/inOutStockService"
	"graduationProjectPeng/service/inventoryService"

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

func StockInventory(c *gin.Context) {
	var param InventoryModel.Inventory
	if err := c.ShouldBindJSON(&param); err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.INVALID_PARAMS, nil)
		return
	}
	if param.Type != InventoryModel.PANYING && param.Type != InventoryModel.LOSE && param.Type != InventoryModel.NORMAL {
		logging.Warn("盘点结果类型参数错误", param.Type)
		common.Json_return(c, e.INVALID_PARAMS, nil)
		return
	}
	if err := inventoryService.AddInventory(&param); err != nil {
		logging.Warn(err.Error(), param)
		common.Json_return(c, e.ERROR, nil)
		return
	}
	logging.Info(param)
	common.Json_return(c, e.SUCCESS, nil)
}

func QueryInventory(c *gin.Context) {
	var param models.InventoryListParam
	if err := c.ShouldBindQuery(&param); err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if param.Type != 0 && param.Type != InventoryModel.PANYING && param.Type != InventoryModel.LOSE && param.Type != InventoryModel.NORMAL {
		logging.Warn("盘点结果类型参数错误", param.Type)
		common.Json_return(c, e.INVALID_PARAMS, nil)
		return
	}
	data, err := inventoryService.QueryInventoryList(&param)
	if err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.ERROR, "")
		return
	}
	common.Json_return(c, e.SUCCESS, data)
}
