package m_dept

import "whacos/models"

type SysDept struct {
	models.Model

	ParentId int    `json:"parentId" gorm:"idx_parent_id"` // 父部门Id
	Name     string `json:"name"`                          // 部门名称
	SortNum  int    `json:"sortNum"`                       // 排序号
}
