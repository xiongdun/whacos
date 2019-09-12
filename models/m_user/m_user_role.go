package m_user

type SysUserRole struct {
	Id     int `json:"id" gorm:"primary_key"` // 主键ID
	UserId int `json:"userId"`                // 用户ID
	RoleId int `json:"roleId"`                // 角色ID
}
