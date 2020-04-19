package InventoryModel

import (
	"graduationProjectPeng/db"
	"graduationProjectPeng/models"
	"time"

	"github.com/jinzhu/gorm"
)

type Inventory struct {
	Id         int   `json:"id" form:"id" gorm:"primary_key"`
	Type       int   `json:"type" binding:"required"`
	GoodsId    int   `json:"goodsId" binding:"required"`
	Num        int64 `json:"num" binding:"required"`
	CategoryId int   `json:"categoryId"`
	CreatedAt  int64
}

type InventoryListRsp struct {
	Inventory
	CategoryName string `json:"categoryName"`
	GoodName     string `json:"goodsName"`
}

const PANYING = 1 //盘盈
const LOSE = 2    //盘亏
const NORMAL = 3  //正常

func (inventory *Inventory) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("createdAt", time.Now().Unix())
}

/**
新增盘点记录
*/
func (Inventory *Inventory) AddInventoryRow() error {
	return db.Db.Create(Inventory).Error
}

func Query(param *models.InventoryListParam) (data []*Inventory, err error) {
	db := db.Db
	if param.CategoryId != 0 {
		db = db.Where("category_id = ?", param.CategoryId)
	}
	if param.GoodsId != 0 {
		db = db.Where("goods_id = ?", param.GoodsId)
	}
	if param.StartAt != 0 {
		db = db.Where("created_at > ?", param.StartAt)
	}
	if param.Type != 0 {
		db = db.Where("type = ?", param.Type)
	}
	if param.Id != 0 {
		db = db.Where("id = ?", param.Id)
	}
	err = db.Find(&data).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}