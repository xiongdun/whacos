package s_file

import (
	"whacos/models"
	"whacos/models/m_file"
)

// 按ID查询
// param id ID
func FindFileById(id int) (*m_file.File, error) {
	var param m_file.File
	return param.SelectById(id)
}

// 查询文件列表
func FindFileList(param m_file.File) ([]m_file.File, error) {
	return param.SelectList(param)
}

// 分页查询
func FindFilePage(param m_file.File, pageNum int, pageSize int) (p models.Page, err error) {
	if dataPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountFile(param); err2 == nil {
			return models.PageUtils(dataPage, count, pageSize), nil
		}
	}
	return
}

// 保存文件
func SaveFile(param m_file.File) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计文件数
func CountFile(param m_file.File) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建文件
func CreateFile(param m_file.File) error {
	return param.Insert(param)
}

// 修改文件
func ModifyFile(param m_file.File) error {
	return param.UpdateById(param)
}

// 按文件ID删除
// param id 文件ID
func RemoveFileById(id int) error {
	file := m_file.File{}
	return file.DeleteById(id)
}
