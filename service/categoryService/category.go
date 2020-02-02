package categoryService

import (
	"graduationProjectPeng/models/categoryModel"
)

/**
获取分类信息
*/
func GetCategory(id, child string) (data []*categoryModel.Cate, err error) {
	dataMap := make(map[int]map[int]*categoryModel.Cate)
	cates := make([]*categoryModel.Cate, 0)
	maps := make(map[string]interface{})
	if id == "" {
		cates, err = categoryModel.GetcateAnd(maps)
	} else {
		maps["id"] = id
		ormaps := make(map[string]interface{})
		if child != "" {
			ormaps["parent"] = id
		}
		cates, err = categoryModel.GetCateOr(maps, ormaps)
	}
	if err != nil {
		return data, err
	}
	dataMap = initCate(cates)
	data = dfsCate(dataMap, 0)
	return data, nil
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
