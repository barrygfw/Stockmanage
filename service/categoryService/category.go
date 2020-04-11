package categoryService

import (
	"graduationProjectPeng/models/categoryModel"
	"strconv"
)

/**
获取分类信息
*/
func GetCategory(idStr, child string) ([]*categoryModel.Cate, error) {
	var index int
	if idStr == "" {
		index = 0
	} else {
		id, _ := strconv.Atoi(idStr)
		if child == "" {
			ids := make([]*int, 0)
			ids = append(ids, &id)
			cates, err := categoryModel.GetCateByIds(ids)
			if err != nil {
				return cates, err
			}
			return cates, nil
		}
		index = id
	}
	return getAllChildCates(index)
}

/**
根据id获取分类及其子分类
id 为0时获取所有分类
*/
func getAllChildCates(parentId int) ([]*categoryModel.Cate, error) {
	allCates, err := categoryModel.GetAllCate()
	if err != nil {
		return allCates, err
	}
	dataMap := initCate(allCates)
	if parentId != 0 {
		ids := make([]*int, 0)
		ids = append(ids, &parentId)
		cates, err := categoryModel.GetCateByIds(ids)
		if err != nil {
			return cates, err
		}
		cates[0].Child = dfsCate(dataMap, parentId)
		return cates, nil
	}
	data := dfsCate(dataMap, parentId)
	return data, nil
}

/**
获取所有分类id
 */
func getCateIds(cates []*categoryModel.Cate) ([]*int, error) {
	ids := make([]*int, 0)
	for _, cate := range cates {
		ids = append(ids, &cate.Id)
		if cate.Child != nil {
			if childIds, err := getCateIds(cate.Child); err == nil {
				ids = append(ids, childIds...)
			}
		}
	}
	return ids, nil
}

/**
添加分类
*/
func AddCategory(cate *categoryModel.Cate) (bool, error) {
	//数据检查
	if ok, err := cate.CheckCategory("insert"); !ok {
		return ok, err
	}
	return categoryModel.AddCate(cate)
}

/**
更新分类
*/
func UpdateCategory(cate *categoryModel.Cate) (bool, error) {
	//数据检查
	if ok, err := cate.CheckCategory("update"); !ok {
		return ok, err
	}
	return categoryModel.UpdateCate(cate)
}

/**
删除分类（支持批量）
*/
func DelCategory(ids []*int) (bool, error) {
	//检查id合法性
	if ok, err := categoryModel.ExistIds(ids); !ok {
		return ok, err
	}
	//检查逻辑合理性（待开发）
	//code......
	return categoryModel.DelCate(ids)
}

/**
结构化category数据
*/
func initCate(cates []*categoryModel.Cate) map[int]map[int]*categoryModel.Cate {
	data := make(map[int]map[int]*categoryModel.Cate)
	if len(cates) < 1 {
		return data
	}
	for _, cate := range cates {
		if _, ok := data[cate.Parent]; !ok {
			data[cate.Parent] = make(map[int]*categoryModel.Cate)
		}
		data[cate.Parent][cate.Id] = cate
	}
	return data
}

/**
深度优先遍历处理所有嵌套子分类
*/
func dfsCate(cate map[int]map[int]*categoryModel.Cate, index int) []*categoryModel.Cate {
	temp := make([]*categoryModel.Cate, 0)
	if _, ok := cate[index]; !ok {
		return temp
	}
	for key, val := range cate[index] {
		if _, ok := cate[key]; ok {
			val.Child = dfsCate(cate, key)
		}
		temp = append(temp, val)
	}
	return temp
}
