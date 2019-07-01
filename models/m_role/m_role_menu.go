package m_role

type SysUserRole struct {
	Id     int `json:"id" gorm:"primary_key"`
	RoleId int `json:"roleId"`
	MenuId int `json:"menuId"`
}
