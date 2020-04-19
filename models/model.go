package models

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type IdList struct {
	Ids []*int `json:"ids" binding:"required"`
}

type InoutListParam struct {
	CategoryId int `json:"categoryId" form:"categoryId"`
	GoodsId int `json:"goodsId" form:"goodsId"`
	StartAt int64 `json:"startAt" form:"startAt"`
	Type    int `json:"type" form:"type"`
	Id      int   `json:"id" form:"id"`
}
