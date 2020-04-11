package category

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/categoryModel"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/service/categoryService"
	"graduationProjectPeng/service/common"

	"github.com/gin-gonic/gin"
)

/**
获取所有分类
无id取所有，有id无child取单独id，有id有child取该id及其子分类
api : /api/category/getall?id=xx&child=xx
params : Query
{
	"id" : 1, //(非必须)
	"child" : "ok" //(非必须)
}
*/
func GetCate(c *gin.Context) {
	data := make([]*categoryModel.Cate, 0)
	var err error
	id := c.Query("id")
	child := c.Query("child")
	data, err = categoryService.GetCategory(id, child)
	if err != nil {
		logging.Warn(err.Error())
		common.Json_return(c, e.ERROR, err.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, data)
}

/**
新增分类
api : /api/category/add
params: json
{
	"parent_id": 111,
	"name": "示例"
}
*/
func AddCate(c *gin.Context) {
	var cate categoryModel.Cate
	if err := c.ShouldBindJSON(&cate); err != nil {
		logging.Warn(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, err.Error())
		return
	}
	if ok, errs := categoryService.AddCategory(&cate); !ok {
		logging.Warn(errs.Error())
		common.Json_return(c, e.ERROR, errs.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

/**
删除分类
支持批量，父级分类删除，所有子分类都会删除
api : /api/category/del
params : json
{
	"ids":[1,2,3,]//id数组
}
*/
func DelCate(c *gin.Context) {
	var ids models.IdList
	if err := c.ShouldBindJSON(&ids); err != nil || len(ids.Ids) < 1 {
		logging.Warn(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, err.Error())
		return
	}
	ok, err := categoryModel.DelCate(ids.Ids)
	if !ok {
		logging.Warn(err.Error())
		common.Json_return(c, e.ERROR, err.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

/**
更新分类
api : api/category/update
params : json
{
	"id" : 1 //需更新的分类id
	"name" : "demo" //更新分类名
	"parent" : 0 //更新分类的父分类
}
*/
func UpdateCate(c *gin.Context) {
	var cate categoryModel.Cate

	if err := c.ShouldBindJSON(&cate); err != nil {
		logging.Warn(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, err.Error())
		return
	}
	if ok, err := categoryService.UpdateCategory(&cate); !ok {
		logging.Warn(err.Error())
		common.Json_return(c, e.ERROR, err.Error())
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}
