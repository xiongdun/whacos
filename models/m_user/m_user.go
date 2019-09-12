package m_user

import (
	"github.com/jinzhu/gorm"
	"time"
	"whacos/models"
)

const tableName = "sys_user user"

type SysUser struct {
	models.Model

	Username    string     `json:"username" gorm:"uk_username"`   // 用户名（唯一）
	RealName    string     `json:"realName"`                      // 真实姓名
	Password    string     `json:"password"`                      // 密码
	DeptId      int        `json:"deptId" gorm:"idx_dept_id"`     // 部门ID
	Email       string     `json:"email" gorm:"uk_email"`         // 邮箱
	MobilePhone string     `json:"mobile" gorm:"uk_mobile_phone"` // 电话号码
	IdCard      string     `json:"idCard" gorm:"uk_id_card"`      // 身份号码
	Status      int        `json:"status"`                        // 用户状态
	Gender      int        `json:"gender"`                        // 性别
	Birth       *time.Time `json:"birth"`                         // 出生日期
	AvatarId    int        `json:"avatarId"`                      // 用户头像 （文件Id）
	WorkAddress string     `json:"workAddress"`                   // 工作详细地址
	Hobby       string     `json:"hobby"`                         // 爱好
	Province    string     `json:"province"`                      // 省
	City        string     `json:"city"`                          // 市
	District    string     `json:"district"`                      // 区
	Remark      string     `json:"remarks"`                       // 备注
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
