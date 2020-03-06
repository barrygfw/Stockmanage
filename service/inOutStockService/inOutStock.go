package inOutStockService

import (
	"graduationProjectPeng/db"
	"graduationProjectPeng/models/goodsModel"
	"graduationProjectPeng/models/inOutStockModel"
	"graduationProjectPeng/pkg/e"
)

/**
商品出入库
*/
func AddInOutStockRow(stock *inOutStockModel.InoutStock) (bool, int) {
	//事务开始
	tx := db.Db.Begin()
	defer tx.Rollback()
	//更新商品库存
	goods := goodsModel.Goods{}
	if err := tx.Where("goods_id = ?", stock.GoodsId).Take(&goods).Error; err != nil {
		return false, e.ERROR_UPDATE_STOCK_FAIL
	}
	if stock.Type == inOutStockModel.InputStock {
		goods.Stock += stock.Num
	} else {
		goods.Stock -= stock.Num
	}
	if err := tx.Save(goods).Error; err != nil {
		return false, e.ERROR_UPDATE_STOCK_FAIL
	}
	if err := tx.Create(stock).Error; err != nil {
		return false, e.ERROR_ADD_INOUT_ROW_FAIL
	}
	tx.Commit()
	return true, e.SUCCESS
}
