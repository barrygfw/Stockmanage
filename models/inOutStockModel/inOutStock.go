package inOutStockModel

import (
	"graduationProjectPeng/db"
	"graduationProjectPeng/models"
	"time"

	"github.com/jinzhu/gorm"
)

type InoutStock struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Type       int    `json:"type" binding:"required"`
	GoodsId    int    `json:"goods_id" binding:"required"`
	Num        int64  `json:"num" binding:"required"`
	CategoryId int    `json:"category_id"`
	Comment    string `json:"comment" binding:"required"`
	CreatedAt  int64
}

type InoutListRsp struct {
	InoutStock
	CategoryName string `json:"categoryName"`
	GoodName     string `json:"goodsName"`
}

const InputStock int = 2  //入库
const OutputStock int = 1 //出库

func (inOutStock *InoutStock) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("createdAt", time.Now().Unix())
}

/**
新增出入库记录
*/
func (inOutStock *InoutStock) AddInoutRow() error {
	return db.Db.Create(inOutStock).Error
}

func Query(param *models.InoutListParam) (data []*InoutStock, err error) {
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
