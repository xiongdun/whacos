package m_user

type SysUserRole struct {
	Id     int `json:"id" gorm:"primary_key"`
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}
