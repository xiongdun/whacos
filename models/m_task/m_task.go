package m_task

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

type Task struct {
	models.Model
	JobName        string `json:"jobName"`
	CronExpression string `json:"cronExpression"`
	ModelName      string `json:"modelName"`
	MethodName     string `json:"methodName"`
	ConcurrentFlag string `json:"concurrentFlag"`
	JobStatus      string `json:"jobStatus"`
	JobGroup       string `json:"jobGroup"`
	Remarks        string `json:"remarks"`
}

// 查询任务记录
func (t *Task) SelectById(id int) (*Task, error) {
	var task Task
	if err := models.DB.Model(&task).Where("id = ? and del_flag = 1", id).Find(&task).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &task, nil
}

// 分页查询列表记录
func (t *Task) SelectPage(param Task, pageNum int, pageSize int) ([]Task, error) {
	var tasks []Task
	if err := models.DB.Where(&param).Order("role.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&tasks).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tasks, nil
}

// 查询任务列表
func (t *Task) SelectList(param Task) ([]Task, error) {
	var tasks []Task
	if err := models.DB.Where(&param).Find(&tasks).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tasks, nil
}

// 统计任务
func (t *Task) Count(param Task) (int, error) {
	var count int
	if err := models.DB.Model(&Task{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除任务
func (t *Task) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增任务
func (t *Task) Insert(task Task) error {
	if err := models.DB.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

// 修改任务
func (t *Task) UpdateById(task Task) error {
	if err := models.DB.Model(&Task{}).Where("id = ?", task.Id).Update(&task).Error; err != nil {
		return err
	}
	return nil
}
