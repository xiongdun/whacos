package s_task

import (
	"whacos/models"
	"whacos/models/m_task"
)

// 按ID查询
// param id ID
func FindTaskById(id int) (*m_task.Task, error) {
	var param m_task.Task
	return param.SelectById(id)
}

// 查询任务列表
func FindTaskList(param m_task.Task) ([]m_task.Task, error) {
	return param.SelectList(param)
}

// 分页查询
func FindTaskPage(param m_task.Task, pageNum int, pageSize int) (p models.Page, err error) {
	if dataPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountTask(param); err2 == nil {
			return models.PageUtils(dataPage, count, pageSize), nil
		}
	}
	return
}

// 保存任务
func SaveTask(param m_task.Task) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计任务数
func CountTask(param m_task.Task) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建任务
func CreateTask(param m_task.Task) error {
	return param.Insert(param)
}

// 修改任务
func ModifyTask(param m_task.Task) error {
	return param.UpdateById(param)
}

// 按任务ID删除
// param id 任务ID
func RemoveTaskById(id int) error {
	task := m_task.Task{}
	return task.DeleteById(id)
}
