package models

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IdList struct {
	Ids []*int `json:"ids" binding:"required"`
}
