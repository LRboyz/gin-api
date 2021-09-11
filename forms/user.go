package forms

type LoginForm struct {
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	UserName string `form:"name" json:"name" binding:"required,mobile"`
}

type UserListForm struct {
	// 页数
	Page int `form:"page" json:"page" binding:"required"`
	// 每页个数
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}
