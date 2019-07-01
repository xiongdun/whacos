package m_user

import (
	"github.com/jinzhu/gorm"
	"time"
	"whacos/models"
)

const tableName = "sys_user user"

type SysUser struct {
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
func (u *SysUser) SelectById(id int) (*SysUser, error) {
	var sysUser SysUser
	if err := models.DB.Model(&sysUser).Where("id = ? and del_flag = 1", id).Find(&sysUser).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &sysUser, nil
}

// 按用户名查询
func (u *SysUser) SelectByUsername(username string) (*SysUser, error) {
	var sysUser SysUser
	if err := models.DB.Model(&sysUser).Where("username = ? and del_flag = 1", username).Find(&sysUser).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &sysUser, nil
}

// 查询用户记录列表
func (u *SysUser) SelectList(param SysUser) ([]SysUser, error) {
	var users []SysUser
	if err := models.DB.Where(&param).Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

// 分页查询列表记录
func (u *SysUser) SelectPage(param SysUser, pageNum int, pageSize int) ([]SysUser, error) {
	var sysUsers []SysUser
	if err := models.DB.Where(&param).Order("user.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&sysUsers).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return sysUsers, nil
}

// 统计记录
func (u *SysUser) Count(param SysUser) (int, error) {
	var count int
	if err := models.DB.Model(&SysUser{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除用户记录
func (u *SysUser) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增记录
func (u *SysUser) Insert(sysUser SysUser) error {
	if err := models.DB.Create(&sysUser).Error; err != nil {
		return err
	}
	return nil
}

// 修改用户信息记录
func (u *SysUser) UpdateById(sysUser SysUser) error {
	if err := models.DB.Model(&SysUser{}).Where("id = ?", sysUser.Id).Update(&sysUser).Error; err != nil {
		return err
	}
	return nil
}
