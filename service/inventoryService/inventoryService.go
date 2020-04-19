package inventoryService

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/InventoryModel"
	"graduationProjectPeng/models/categoryModel"
	"graduationProjectPeng/models/goodsModel"
)

func AddInventory(param *InventoryModel.Inventory) error {
	goodsId := []*int{&param.GoodsId}
	goodsInfo, _ := goodsModel.GetGoodsByIds(goodsId)
	param.CategoryId = goodsInfo[0].CategoryId
	return param.AddInventoryRow()
}

func QueryInventoryList(param *models.InventoryListParam) ([]*InventoryModel.InventoryListRsp, error) {
	data, err := InventoryModel.Query(param)
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
	resData := make([]*InventoryModel.InventoryListRsp, 0)
	for _, inventory := range data {
		resData = append(resData, &InventoryModel.InventoryListRsp{
			Inventory:    *inventory,
			CategoryName: cateIdNameMap[inventory.CategoryId],
			GoodName:     goodsIdNameMap[inventory.GoodsId],
		})
	}
	return resData, nil
}
