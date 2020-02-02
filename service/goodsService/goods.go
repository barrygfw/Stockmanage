package goodsService

import "graduationProjectPeng/models/goodsModel"

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
