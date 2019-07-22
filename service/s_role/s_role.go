package s_role

import (
	"whacos/models/m_role"
)

// 按ID查询
// param id ID
func FindRoleById(id int) (*m_role.SysRole, error) {
	var param m_role.SysRole
	return param.SelectById(id)
}

// 查询角色列表
func FindRoleList(param m_role.SysRole) ([]m_role.SysRole, error) {
	return param.SelectList(param)
}
