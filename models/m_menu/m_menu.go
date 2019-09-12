package m_menu

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

type SysMenu struct {
	models.Model

	ParentId    int    `json:"parentId" gorm:"idx_parent_id"` // 父菜单ID
	Name        string `json:"name"`                          // 菜单名称
	Url         string `json:"url"`                           // 菜单url
	Permissions string `json:"permissions"`                   // 菜单权限
	MenuType    int    `json:"menuType"`                      // 菜单类型
	Icon        string `json:"icon"`                          // 菜单图标
	OrderNum    int    `json:"orderNum"`                      // 排序号
	Remark      string `json:"remark"`                        // 备注
}

// 查询指定菜单记录
func (m *SysMenu) SelectById(id int) (*SysMenu, error) {
	var sysMenu SysMenu
	if err := models.DB.Model(&sysMenu).Where("id = ? and del_flag = 1", id).Find(&sysMenu).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &sysMenu, nil
}

// 分页查询列表记录
func (m *SysMenu) SelectPage(param SysMenu, pageNum int, pageSize int) ([]SysMenu, error) {
	var sysMenus []SysMenu
	if err := models.DB.Where(&param).Order("menu.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&sysMenus).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return sysMenus, nil
}

// 查询菜单记录列表
func (m *SysMenu) SelectList(param SysMenu) ([]SysMenu, error) {
	var menus []SysMenu
	if err := models.DB.Where(&param).Find(&menus).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return menus, nil
}

// 统计菜单记录
func (m *SysMenu) Count(param SysMenu) (int, error) {
	var count int
	if err := models.DB.Model(&SysMenu{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除菜单记录
func (m *SysMenu) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增菜单记录
func (m *SysMenu) Insert(sysMenu SysMenu) error {
	if err := models.DB.Create(&sysMenu).Error; err != nil {
		return err
	}
	return nil
}

// 修改菜单记录
func (m *SysMenu) UpdateById(sysMenu SysMenu) error {
	if err := models.DB.Model(&SysMenu{}).Where("id = ?", sysMenu.Id).Update(&sysMenu).Error; err != nil {
		return err
	}
	return nil
}
