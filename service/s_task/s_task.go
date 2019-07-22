package s_task

import (
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
