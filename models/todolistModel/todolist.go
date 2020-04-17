package todolistModel

import (
	"graduationProjectPeng/db"
	"time"
)

type Todolist struct {
	Id        int    `json:"id" gorm:"primary_key"`
	StartAt   int64  `json:"start_at" binding:"required"`
	EndAt     int64  `json:"end_at" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CreatedAt int64
	UpdatedAt int64
	Status    int `json:"status"`
}

//func (todolist *Todolist) BeforeCreate(scope gorm.Scope) error {
//	return scope.SetColumn("createdAt", time.Now().Unix())
//}

//func (todolist *Todolist) BeforeSave(scope gorm.Scope) error {
//	return scope.SetColumn("updatedAt", time.Now().Unix())
//}

func (Todo *Todolist) AddToDo() error {
	Todo.CreatedAt = time.Now().Unix()
	Todo.UpdatedAt = time.Now().Unix()
	Todo.Status = 2
	return db.Db.Create(Todo).Error
}

func (Todo *Todolist) UptToDo() error {
	Todo.UpdatedAt = time.Now().Unix()
	return db.Db.Save(Todo).Error
}

func DelToDoLists(ids []*int) error {
	return db.Db.Where("id in (?)", ids).Delete(&Todolist{}).Error
}

func Query(where map[string]string) (Todolist []*Todolist, err error) {
	db := db.Db
	if _, ok := where["start_at"]; ok {
		db = db.Where("start_at > ?", where["start_at"])
	}
	if _, ok := where["end_at"]; ok {
		db = db.Where("end_at < ?", where["end_at"])
	}
	if _, ok := where["status"]; ok {
		db = db.Where("status = ?", where["status"])
	}
	err = db.Find(&Todolist).Error
	return
}
