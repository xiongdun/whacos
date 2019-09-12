package m_task

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

type Task struct {
	models.Model
	TaskName       string `json:"taskName" gorm:"idx_task_name"` // 任务名称
	CronExpression string `json:"cronExpression"`                // cron 表达式
	ModelName      string `json:"modelName"`                     // 模块名称
	MethodName     string `json:"methodName"`                    // 方法名称
	ConcurrentFlag string `json:"concurrentFlag"`                // 并发标记
	Status         string `json:"status"`                        // 任务状态
	TaskGroup      string `json:"taskGroup"`                     // 任务组
	Remark         string `json:"remarks"`                       // 备注
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
