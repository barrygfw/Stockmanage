package goodsModel

import (
	"errors"
	"fmt"
	"graduationProjectPeng/db"
	"graduationProjectPeng/models/categoryModel"
	"graduationProjectPeng/models/inOutStockModel"
	"time"

	"github.com/jinzhu/gorm"
)

type Goods struct {
	GoodsId    int    `json:"goods_id" gorm:"primary_key"`
	Name       string `json:"name" binding:"required"`
	CategoryId int    `json:"category_id" binding:"required"`
	CategoryName string `json:"category_name"`
	Icon       string `json:"icon"`
	CreatedAt  int64
	UpdatedAt  int64
	Stock      int64 `json:"stock"`
}

func (goods *Goods) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("createdAt", time.Now().Unix())
}

func QueryGoods(where map[string]string) (goods []*Goods, err error) {
	db := db.Db
	if _, ok := where["category_id"]; ok {
		db = db.Where("category_id = ?", where["category_id"])
	}
	if _, ok := where["name"]; ok {
		db = db.Where("name LIKE ?", "%"+where["name"]+"%")
	}
	err = db.Find(&goods).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

/**
新增商品
*/
func (goods *Goods) AddGoods() error {
	return db.Db.Create(goods).Error
}

/**
更新商品基本信息
*/
func (goods *Goods) UpdateGoods() error {
	goods.UpdatedAt = time.Now().Unix()
	return db.Db.Model(&Goods{}).Omit("stock").UpdateColumns(goods).Error
}

/**
更新商品库存
*/
func UpdateStock(goodsId, inOutType int, num int64) error {
	goods := Goods{}
	if err := db.Db.Where("goods_id = ?", goodsId).Take(&goods).Error; err != nil {
		return err
	}
	if inOutType == inOutStockModel.InputStock {
		goods.Stock += num
	} else {
		goods.Stock -= num
	}
	return db.Db.Save(goods).Error
}

/**
删除商品
*/
func DelGoods(goodsIds []*int) error {
	return db.Db.Where("goods_id in (?)", goodsIds).Delete(&Goods{}).Error
}

/**
相关数据检查
*/
func (goods *Goods) CheckGoods(option string) error {
	switch option {
	case "update":
		ids := make([]*int, 0)
		ids = append(ids, &goods.GoodsId)
		if err := ExistIds(ids); err != nil {
			return err
		}
		fallthrough
	case "insert":
		if err := existGoods(goods.CategoryId, goods.Name); err != nil {
			return err
		}
		cateIds := make([]*int, 0)
		cateIds = append(cateIds, &goods.CategoryId)
		if _, err := categoryModel.ExistIds(cateIds); err != nil {
			return err
		}
	}
	return nil
}

/**
检查该商品是否已经存在
*/
func existGoods(category_id int, name string) error {
	n := 0
	if err := db.Db.Model(&Goods{}).Where("category_id = ? and name = ?", category_id, name).Count(&n).Error; err != nil {
		return err
	}
	if n != 0 {
		return errors.New("该商品已存在")
	}
	return nil
}

/**
检查id是否存在
*/
func ExistIds(ids []*int) error {
	var goods []*Goods
	if err := db.Db.Select("goods_id, name").Where("goods_id in (?)", ids).Find(&goods).Error; err != nil {
		return err
	}
	if len(goods) < 1 {
		return errors.New("请求的商品id不存在")
	}
	goodsMap := make(map[int]string)
	for _, good := range goods {
		goodsMap[good.GoodsId] = good.Name
	}
	for _, id := range ids {
		if _, ok := goodsMap[*id]; !ok {
			return errors.New(fmt.Sprintf("商品id : %d 不存在", *id))
		}
	}
	return nil
}
