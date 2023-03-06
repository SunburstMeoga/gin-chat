package service

import (
	"gochat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 用户列表
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
// @Summary 新增用户
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

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code", "删除成功"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code", "修改成功"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")

	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改成功",
	})
}
