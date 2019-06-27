package m_user

import (
	"github.com/jinzhu/gorm"
	"time"
	"whacos/models"
)

const tableName = "sys_user user"

type User struct {
	models.Model

	Username    string    `json:"username" gorm:"idx_username"`
	Name        string    `json:"name" gorm:"idx_name"`
	Password    string    `json:"password"`
	DeptId      int       `json:"deptId" gorm:"idx_dept_id"`
	Email       string    `json:"email" gorm:"idx_email"`
	Mobile      string    `json:"mobile" gorm:"idx_mobile"`
	IdCard      string    `json:"idCard" gorm:"idx_id_card"`
	Status      int       `json:"status"`
	Sex         int       `json:"sex"`
	Birth       time.Time `json:"birth"`
	PicId       int       `json:"picId"`
	LiveAddress string    `json:"liveAddress"`
	Hobby       string    `json:"hobby"`
	Province    string    `json:"province"`
	City        string    `json:"city"`
	District    string    `json:"district"`
	Remarks     string    `json:"remarks"`
}

// 查询指定用户记录
func (u *User) SelectById(id int) (*User, error) {
	var user User
	if err := models.DB.Model(&user).Table(tableName).Where("id = ?", id).Find(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

// 查询用户记录列表
func (u *User) SelectList(param User) ([]User, error) {
	var users []User
	if err := models.DB.Table(tableName).Where(&param).Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

// 分页查询列表记录
func (u *User) SelectPage(param User, pageNum int, pageSize int) ([]User, error) {
	var users []User
	if err := models.DB.Table(tableName).Where(&param).Order("user.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

// 统计记录
func (u *User) Count(param User) (int, error) {
	var count int
	if err := models.DB.Model(&User{}).Table(tableName).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除用户记录
func (u *User) DeleteById(id int) error {
	if err := models.DB.Table("sys_user").Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}
	return nil
}

// 新增记录
func (u *User) Insert(user User) error {
	if err := models.DB.Table("sys_user").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 修改用户信息记录
func (u *User) UpdateById(user User) error {
	if err := models.DB.Model(&User{}).Table("sys_user").Where("id = ?", user.Id).Update(&user).Error; err != nil {
		return err
	}
	return nil
}
