package s_dict

import (
	"whacos/models/m_dict"
)

// 按ID查询
// param id ID
func FindDictById(id int) (*m_dict.Dict, error) {
	var param m_dict.Dict
	return param.SelectById(id)
}

// 查询字典列表
func FindDictList(param m_dict.Dict) ([]m_dict.Dict, error) {
	return param.SelectList(param)
}
