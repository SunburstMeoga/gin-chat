package models

import (
	"fmt"
	"gochat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string `valid: "email"`
	Identity      string
	Salt          string
	ClientIp      string
	ClientPort    string
	LoginTime     string
	HeartbeatTime string
	LoginOutTime  string
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// 通过name和password查找user
func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	utils.DB.Where("name = ? and pass_word = ? ", name, password).First(&user)
	return user
}

// 通过phone查找user
func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Phone = ?", phone).First(&user)
}

// 通过email查找user
func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Email = ?", email).First(&user)
}

// 通过name查找user
func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
