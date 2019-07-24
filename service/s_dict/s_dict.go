package s_dict

import (
	"whacos/models"
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

// 分页查询
func FindDictPage(param m_dict.Dict, pageNum int, pageSize int) (p models.Page, err error) {
	if dataPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountDict(param); err2 == nil {
			return models.PageUtils(dataPage, count, pageSize), nil
		}
	}
	return
}

// 保存字典
func SaveDict(param m_dict.Dict) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计字典数
func CountDict(param m_dict.Dict) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建字典
func CreateDict(param m_dict.Dict) error {
	return param.Insert(param)
}

// 修改字典
func ModifyDict(param m_dict.Dict) error {
	return param.UpdateById(param)
}

// 按字典ID删除
// param id 字典ID
func RemoveDictById(id int) error {
	dict := m_dict.Dict{}
	return dict.DeleteById(id)
}
