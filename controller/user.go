package controller

import (
	"github.com/LRboyz/gin-api/Response"
	"github.com/LRboyz/gin-api/dao"
	"github.com/LRboyz/gin-api/forms"
	"github.com/LRboyz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

// 登录
func PasswordLogin(c *gin.Context) {
	loginForm := forms.LoginForm{}
	if err := c.ShouldBind(&loginForm); err != nil {
		// 统一异常处理
		utils.HandleValidatorError(c, err)
		return
	}
	Response.Success(c, 200, "success", "test")
}

func GetUserList(c *gin.Context) {
	// 获取参数
	UserListForm := forms.UserListForm{}
	if err := c.ShouldBind(&UserListForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	// 获取数据
	total, userlist := dao.GetUserListDao(UserListForm.Page, UserListForm.PageSize)
	if (total + len(userlist)) == 0 {
		Response.Err(c, 400, 400, "未获取到到数据", map[string]interface{}{
			"total":    total,
			"userlist": userlist,
		})
		return
	}
	Response.Success(c, 200, "获取用户列表成功", map[string]interface{}{
		"total":    total,
		"userlist": userlist,
	})

}
