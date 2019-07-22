package s_log

import (
	"whacos/models/m_log"
)

// 按ID查询
// param id ID
func FindLogById(id int) (*m_log.Log, error) {
	var param m_log.Log
	return param.SelectById(id)
}

// 查询日志列表
func FindLogList(param m_log.Log) ([]m_log.Log, error) {
	return param.SelectList(param)
}
