package m_log

import "whacos/models"

type Log struct {
	models.Model
	UserId    int    `json:"userId" gorm:"idx_user_id"`
	Username  string `json:"username"`
	Operation string `json:"operation"`
	Time      int    `json:"time"`
	Method    string `json:"method"`
	Params    string `json:"params"`
	Ip        string `json:"ip"`
}
