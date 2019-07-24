package m_role

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

type SysRole struct {
	models.Model

	Name    string `json:"name"`
	Code    string `json:"code"`
	Remarks string `json:"remarks"`
}

// 查询指定用户记录
func (r *SysRole) SelectById(id int) (*SysRole, error) {
	var sysRole SysRole
	if err := models.DB.Model(&sysRole).Where("id = ? and del_flag = 1", id).Find(&sysRole).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &sysRole, nil
}

// 分页查询列表记录
func (r *SysRole) SelectPage(param SysRole, pageNum int, pageSize int) ([]SysRole, error) {
	var sysRoles []SysRole
	if err := models.DB.Where(&param).Order("role.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&sysRoles).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return sysRoles, nil
}

// 查询角色记录列表
func (r *SysRole) SelectList(param SysRole) ([]SysRole, error) {
	var roles []SysRole
	if err := models.DB.Where(&param).Find(&roles).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roles, nil
}

// 统计记录
func (r *SysRole) Count(param SysRole) (int, error) {
	var count int
	if err := models.DB.Model(&SysRole{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除角色记录
func (r *SysRole) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增记录
func (r *SysRole) Insert(sysRole SysRole) error {
	if err := models.DB.Create(&sysRole).Error; err != nil {
		return err
	}
	return nil
}

// 修改角色记录
func (r *SysRole) UpdateById(sysRole SysRole) error {
	if err := models.DB.Model(&SysRole{}).Where("id = ?", sysRole.Id).Update(&sysRole).Error; err != nil {
		return err
	}
	return nil
}
