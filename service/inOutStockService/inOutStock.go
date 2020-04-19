package inOutStockService

import (
	"graduationProjectPeng/db"
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/categoryModel"
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
	stock.CategoryId = goods.CategoryId
	if err := tx.Create(stock).Error; err != nil {
		return false, e.ERROR_ADD_INOUT_ROW_FAIL
	}
	tx.Commit()
	return true, e.SUCCESS
}

func QueryInoutStockList(param *models.InoutListParam) ([]*inOutStockModel.InoutListRsp, error) {
	data, err := inOutStockModel.Query(param)
	if err != nil {
		return nil, err
	}
	goodsIds := make([]*int, 0)
	categoryIds := make([]*int, 0)
	for _, inout := range data {
		goodsIds = append(goodsIds, &inout.GoodsId)
		categoryIds = append(categoryIds, &inout.CategoryId)
	}
	categorysInfo, _ := categoryModel.GetCateByIds(categoryIds)
	goodsInfo, _ := goodsModel.GetGoodsByIds(goodsIds)
	goodsIdNameMap := make(map[int]string)
	cateIdNameMap := make(map[int]string)
	for _, cate := range categorysInfo {
		cateIdNameMap[cate.Id] = cate.Name
	}
	for _, good := range goodsInfo {
		goodsIdNameMap[good.GoodsId] = good.Name
	}
	resData := make([]*inOutStockModel.InoutListRsp, 0)
	for _, inout := range data {
		resData = append(resData, &inOutStockModel.InoutListRsp{
			InoutStock:   *inout,
			CategoryName: cateIdNameMap[inout.CategoryId],
			GoodName:     goodsIdNameMap[inout.GoodsId],
		})
	}
	return resData, nil
}
