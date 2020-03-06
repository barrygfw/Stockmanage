package inOutStockModel

import (
	"graduationProjectPeng/db"
	"time"

	"github.com/jinzhu/gorm"
)

type InoutStock struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Type       int    `json:"type" binding:"required"`
	GoodsId    int    `json:"goods_id" binding:"required"`
	Num        int64  `json:"num" binding:"required"`
	CategoryId int    `json:"category_id" binding:"required"`
	Comment    string `json:"comment" binding:"required"`
	CreatedAt  int64
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
