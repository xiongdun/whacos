package m_log

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

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

// 查询日志记录
func (l *Log) SelectById(id int) (*Log, error) {
	var log Log
	if err := models.DB.Model(&log).Where("id = ? and del_flag = 1", id).Find(&log).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &log, nil
}

// 分页查询列表记录
func (l *Log) SelectPage(param Log, pageNum int, pageSize int) ([]Log, error) {
	var logs []Log
	if err := models.DB.Where(&param).Order("log.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&logs).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return logs, nil
}

// 查询日志列表
func (l *Log) SelectList(param Log) ([]Log, error) {
	var logs []Log
	if err := models.DB.Where(&param).Find(&logs).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return logs, nil
}

// 统计日志
func (l *Log) Count(param Log) (int, error) {
	var count int
	if err := models.DB.Model(&Log{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除日志
func (l *Log) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增日志
func (l *Log) Insert(log Log) error {
	if err := models.DB.Create(&log).Error; err != nil {
		return err
	}
	return nil
}

// 修改日志
func (l *Log) UpdateById(log Log) error {
	if err := models.DB.Model(&Log{}).Where("id = ?", log.Id).Update(&log).Error; err != nil {
		return err
	}
	return nil
}
