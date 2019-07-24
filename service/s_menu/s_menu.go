package s_menu

import (
	"whacos/models"
	"whacos/models/m_menu"
)

// 按ID查询
// param id ID
func FindMenuById(id int) (*m_menu.SysMenu, error) {
	var param m_menu.SysMenu
	return param.SelectById(id)
}

// 查询菜单列表
func FindMenuList(param m_menu.SysMenu) ([]m_menu.SysMenu, error) {
	return param.SelectList(param)
}

// 分页查询
func FindMenuPage(param m_menu.SysMenu, pageNum int, pageSize int) (p models.Page, err error) {
	if dataPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountMenu(param); err2 == nil {
			return models.PageUtils(dataPage, count, pageSize), nil
		}
	}
	return
}

// 保存菜单
func SaveMenu(param m_menu.SysMenu) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计菜单数
func CountMenu(param m_menu.SysMenu) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建菜单
func CreateMenu(param m_menu.SysMenu) error {
	return param.Insert(param)
}

// 修改菜单
func ModifyMenu(param m_menu.SysMenu) error {
	return param.UpdateById(param)
}

// 按菜单ID删除
// param id 菜单ID
func RemoveMenuById(id int) error {
	sysMenu := m_menu.SysMenu{}
	return sysMenu.DeleteById(id)
}
