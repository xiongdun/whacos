package m_log

import (
	"github.com/jinzhu/gorm"
	"time"
	"whacos/models"
)

type Log struct {
	Id          int        `json:"id" gorm:"primary_key"`     // 主键ID
	UserId      int        `json:"userId" gorm:"idx_user_id"` // 用户Id
	Username    string     `json:"username"`                  // 用户名
	Operation   string     `json:"operation"`                 // 操作类型
	Times       int        `json:"times"`                     // 耗费时长
	Method      string     `json:"method"`                    // 请求方法
	Params      string     `json:"params"`                    // 请求参数
	IpAddress   string     `json:"ipAddress"`                 // ip地址
	CreatedBy   int        `json:"createdBy"`                 // 创建人
	CreatedTime *time.Time `json:"createdTime"`               // 创建时间
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
