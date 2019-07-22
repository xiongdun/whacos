package s_file

import (
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
