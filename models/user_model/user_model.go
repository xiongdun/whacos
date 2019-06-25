package user_model

import (
	"time"
	"whacos/models"
)

type User struct {
	models.Model

	Username    string    `json:"username" gorm:"username"`
	Name        string    `json:"name" gorm:"name"`
	Password    string    `json:"password" gorm:"password"`
	DeptId      int       `json:"deptId" gorm:"dept_id"`
	Email       string    `json:"email" gorm:"email"`
	Mobile      string    `json:"mobile" gorm:"mobile"`
	IdCard      string    `json:"idCard" gorm:"id_card"`
	Status      int       `json:"status" gorm:"status"`
	Sex         int       `json:"sex" gorm:"sex"`
	Birth       time.Time `json:"birth" gorm:"birth"`
	PicId       int       `json:"picId" gorm:"pic_id"`
	LiveAddress string    `json:"liveAddress" gorm:"live_address"`
	Hobby       string    `json:"hobby" gorm:"hobby"`
	Province    string    `json:"province" gorm:"province"`
	City        string    `json:"city" gorm:"city"`
	District    string    `json:"district" gorm:"district"`
	Remarks     string    `json:"remarks" gorm:"remarks"`
}

func SelectUserById(id int) (user User) {
	models.DB.Model(&user).Table("sys_user").Where("id = ?", id).Find(&user)
	return
}

func SelectUserList(pageNum int, pageSize int, maps interface{}) (users []User) {
	models.DB.Table("sys_user").Where(&maps).Offset(pageNum).Limit(pageSize).Find(&users)
	return
}
func CountUser(maps interface{}) (count int) {
	models.DB.Model(&User{}).Table("sys_user").Where(&maps).Count(&count)
	return
}

func DeleteUserById(id int) bool {
	models.DB.Table("sys_user").Where("id = ?", id).Delete(&User{})
	return true
}

func InsertUser(user *User) bool {
	models.DB.Table("sys_user").Create(&user)
	return true
}

func UpdateUserById(user *User) bool {
	models.DB.Model(&User{}).Where("id = ?", user.Id).Update(&user)
	return true
}
