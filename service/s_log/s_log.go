package s_log

import (
	"whacos/models"
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

// 分页查询
func FindLogPage(param m_log.Log, pageNum int, pageSize int) (p models.Page, err error) {
	if dataPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountLog(param); err2 == nil {
			return models.PageUtils(dataPage, count, pageSize), nil
		}
	}
	return
}

// 保存日志
func SaveLog(param m_log.Log) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计日志数
func CountLog(param m_log.Log) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建日志
func CreateLog(param m_log.Log) error {
	return param.Insert(param)
}

// 修改日志
func ModifyLog(param m_log.Log) error {
	return param.UpdateById(param)
}

// 按日志ID删除
// param id 日志ID
func RemoveLogById(id int) error {
	sysLog := m_log.Log{}
	return sysLog.DeleteById(id)
}
