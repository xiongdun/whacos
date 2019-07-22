package s_menu

import (
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
