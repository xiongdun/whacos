package s_role

import (
	"whacos/models"
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

// 分页查询
func FindRolePage(param m_role.SysRole, pageNum int, pageSize int) (p models.Page, err error) {
	if dataPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountRole(param); err2 == nil {
			return models.PageUtils(dataPage, count, pageSize), nil
		}
	}
	return
}

// 保存角色
func SaveRole(param m_role.SysRole) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计角色数
func CountRole(param m_role.SysRole) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建角色
func CreateRole(param m_role.SysRole) error {
	return param.Insert(param)
}

// 修改角色
func ModifyRole(param m_role.SysRole) error {
	return param.UpdateById(param)
}

// 按角色ID删除
// param id 角色ID
func RemoveRoleById(id int) error {
	sysRole := m_role.SysRole{}
	return sysRole.DeleteById(id)
}
