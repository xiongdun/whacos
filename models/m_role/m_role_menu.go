package m_role

type SysUserRole struct {
	Id     int `json:"id" gorm:"primary_key"` // 主键Id
	RoleId int `json:"roleId"`                // 角色ID
	MenuId int `json:"menuId"`                // 菜单ID
}
