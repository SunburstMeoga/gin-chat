package service

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
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
// @param name query string false "name"
// @param password query string false "password"
// @param repassword query string false "repassword"
// @Success 200 {string} json{"code", "message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"message": "用户名已存在",
		})
		return
	}

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码输入不一致",
		})
		return
	}
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "注册成功",
	})

}

// FindUserByNameAndPwd
// @Summary 通过用户名和密码查找用户
// @Tags 用户模块
// @param name query string false "name"
// @param password query string false "password"
// @Success 200 {string} json{"code", "message"}
// @Router /user/FindUserByNameAndPwd [get]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "该用户不存在",
		})
		return
	}
	fmt.Println("user-----", user)
	flag := utils.ValidPassword(password, name, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"message": "用户名或密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"message": data,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @param name query string false "name"
// @param password query string false "password"
// @Success 200 {string} json{"code", "message"}
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
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} json{"code", "message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")

	_, err := govalidator.ValidateStruct(user)
	fmt.Println("err------", err)
	if err != nil {
		c.JSON(-1, gin.H{
			"message": "参数不匹配",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改用户信息成功",
	})
}
