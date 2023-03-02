package service

import (
	"gochat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Sunmmary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Sunmmary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repasword := c.Query("repassword")
	if password != repasword {
		c.JSON(-1, gin.H{
			"message": "两次密码输入不一致",
		})
	} else {
		user.PassWord = password
		models.CreateUser(user)
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	}

}
