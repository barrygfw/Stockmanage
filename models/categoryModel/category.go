package categoryModel

import (
	"errors"
	"fmt"
	"graduationProjectPeng/db"

	"github.com/jinzhu/gorm"
)

type Cate struct {
	Id     int     `json:"category_id" gorm:"primary_key"`
	Parent int     `json:"parent_id"`
	Name   string  `json:"name" binding:"required"`
	Child  []*Cate `json:"child_cate"`
}

/**
And 条件查询category
*/
func GetcateAnd(maps interface{}) (cate []*Cate, err error) {
	err = db.Db.Where(maps).Find(&cate).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

/**
Or 条件查询category
*/
func GetCateOr(maps interface{}, ormaps interface{}) (cate []*Cate, err error) {
	err = db.Db.Where(maps).Or(ormaps).Find(&cate).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

/**
插入category
*/
func AddCate(cate *Cate) (bool, error) {
	if err := db.Db.Create(cate).Error; err != nil {
		return false, err
	}
	return true, nil
}

/**
删除分类
父级分类删除子分类跟随删除
*/
func DelCate(ids []*int) (bool, error) {
	//开启事务
	tx := db.Db.Begin()
	if err := tx.Where("id in (?)", ids).Delete(&Cate{}).Error; err != nil {
		//父分类删除失败，回滚
		tx.Rollback()
		return false, err
	}
	if err := tx.Where("parent in (?)", ids).Delete(&Cate{}).Error; err != nil {
		//子分类删除失败，回滚
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}

/**
更新category
*/
func UpdateCate(cate *Cate) (bool, error) {
	if err := db.Db.Model(&Cate{}).Updates(cate).Error; err != nil {
		return false, err
	}
	return true, nil
}

/**
操作前的相关数据检查
*/
func (cate *Cate) CheckCategory(option string) (bool, error) {
	switch option {
	case "update":
		ids := make([]*int, 0)
		ids = append(ids, &cate.Id)
		if ok, err := ExistIds(ids); !ok {
			return ok, err
		}
		fallthrough
	case "insert":
		if ok, err := existParent(cate.Parent); !ok {
			return ok, err
		}
		if ok, err := existName(cate.Parent, cate.Name); !ok {
			return ok, err
		}
	}
	return true, nil
}

/**
检查该分类名称是否存在
*/
func existName(parent int, name string) (bool, error) {
	n := 0
	if err := db.Db.Model(&Cate{}).Where("parent = ? and name = ?", parent, name).Count(&n).Error; err != nil {
		return false, err
	}
	if n != 0 {
		return false, errors.New("该分类名称已存在")
	}
	return true, nil
}

/**
检查父级分类是否存在
*/
func existParent(parent int) (bool, error) {
	n := 0
	if err := db.Db.Model(&Cate{}).Where("id = ?", parent).Count(&n).Error; err != nil {
		return false, err
	}
	if n == 0 {
		return false, errors.New("父级分类不存在")
	}
	return true, nil
}

/**
检查该分类id是否存在
*/
func ExistIds(ids []*int) (bool, error) {
	cates := make([]*Cate, 0)
	if err := db.Db.Select("id, name").Where("id in (?)", ids).Find(&cates).Error; err != nil {
		return false, err
	}
	if len(cates) < 1 {
		return false, errors.New(fmt.Sprintf("请求的分类id不存在"))
	}
	catesMap := make(map[int]string)
	for _, cate := range cates {
		catesMap[cate.Id] = cate.Name
	}
	for _, id := range ids {
		if _, ok := catesMap[*id]; !ok {
			return false, errors.New(fmt.Sprintf("分类id : %d 不存在", *id))
		}
	}
	return true, nil
}

/**
待开发......
删除时检查该分类下是否存在商品
*/
func IdExistGoods(ids []*int) (bool, error) {
	return true, nil
}
