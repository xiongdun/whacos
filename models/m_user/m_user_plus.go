package m_user

import "whacos/models"

type SysUserPlus struct {
	models.Model
	UserId   int    `json:"userId" gorm:"idx_user_id"`
	BankCard string `json:"bankCard"`
	BankName string `json:"bankName"`
}
