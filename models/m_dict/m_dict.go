package m_dict

import "whacos/models"

type Dict struct {
	models.Model

	Name     string `json:"name"`
	Value    string `json:"value"`
	DictType string `json:"dictType"`
	SortNum  int    `json:"sortNum"`
	ParentId int    `json:"parentId" gorm:"idx_parent_id"`
}
