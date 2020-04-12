package goodsService

import (
	"graduationProjectPeng/models/categoryModel"
	"graduationProjectPeng/models/goodsModel"
)

/**
新增商品
*/
func InsertGoods(goods *goodsModel.Goods) error {
	//数据检查
	if err := goods.CheckGoods("insert"); err != nil {
		return err
	}
	return goods.AddGoods()
}

/**
更新商品
*/
func UpdateGoods(goods *goodsModel.Goods) error {
	//数据检查
	if err := goods.CheckGoods("update"); err != nil {
		return err
	}
	return goods.UpdateGoods()
}

/**
删除商品
*/
func DeleteGoods(goodsIds []*int) error {
	//数据检查
	if err := goodsModel.ExistIds(goodsIds); err != nil {
		return err
	}
	return goodsModel.DelGoods(goodsIds)
}

/**
查询商品
*/
func QueryGoods(goodsName, categoryId string) ([]*goodsModel.Goods, error) {
	where := make(map[string]string)
	if categoryId != "" {
		where["category_id"] = categoryId
	}
	if goodsName != "" {
		where["name"] = goodsName
	}
	goodsList, err := goodsModel.QueryGoods(where)
	if err != nil {
		return goodsList, err
	}
	cateIdList := make([]*int, 0)
	for _, goods := range goodsList {
		cateIdList = append(cateIdList, &goods.CategoryId)
	}
	catesList, err := categoryModel.GetCateByIds(cateIdList)
	if err != nil {
		return goodsList, err
	}
	cateIdNameMap := make(map[int]string, 0)
	for _, cate := range catesList {
		cateIdNameMap[cate.Id] = cate.Name
	}
	for index, _ := range goodsList {
		goodsList[index].CategoryName = cateIdNameMap[goodsList[index].CategoryId]
	}
	return goodsList, nil
}
