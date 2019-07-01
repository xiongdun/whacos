package m_user

import "whacos/models"

type SysUserPlus struct {
	models.Model
	userId   int    `json:"userId" gorm:"idex_user_id"`
	bankCard string `json:"bankCard"`
	bankName string `json:"bankName"`
	password string `json:"password"`
}
