package toDoList

import (
	"graduationProjectPeng/models"
	"graduationProjectPeng/models/todolistModel"
	"graduationProjectPeng/pkg/e"
	"graduationProjectPeng/pkg/logging"
	"graduationProjectPeng/service/common"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var todo todolistModel.Todolist
	if err := c.ShouldBindJSON(&todo); err != nil {
		logging.Warn(c.Params, e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if err := todo.AddToDo(); err != nil {
		logging.Warn(e.GetMsg(e.ERROR), err.Error())
		common.Json_return(c, e.ERROR, "")
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

func UptTodo(c *gin.Context) {
	var todo todolistModel.Todolist
	if err := c.ShouldBindJSON(&todo); err != nil {
		logging.Warn(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if err := todo.UptToDo(); err != nil {
		logging.Warn(e.GetMsg(e.ERROR), err.Error())
		common.Json_return(c, e.ERROR, "")
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

func DelTodo(c *gin.Context) {
	var ids models.IdList
	if err := c.ShouldBindJSON(&ids); err != nil {
		logging.Warn(e.GetMsg(e.INVALID_PARAMS), err.Error())
		common.Json_return(c, e.INVALID_PARAMS, "")
		return
	}
	if err := todolistModel.DelToDoLists(ids.Ids); err != nil {
		logging.Warn(e.GetMsg(e.ERROR), err.Error())
		common.Json_return(c, e.ERROR, "")
		return
	}
	common.Json_return(c, e.SUCCESS, "")
}

func Query(c *gin.Context) {
	startAt := c.Query("start_at")
	endAt := c.Query("end_at")
	status := c.Query("status")
	where := make(map[string]string)
	if startAt != "" {
		where["start_at"] = startAt
	}
	if endAt != "" {
		where["end_at"] = endAt
	}
	if status != "" {
		where["status"] = status
	}
	data, err := todolistModel.Query(where)
	if err != nil {
		logging.Warn(e.GetMsg(e.ERROR), err.Error())
		common.Json_return(c, e.ERROR, "")
		return
	}
	common.Json_return(c, e.SUCCESS, data)
}