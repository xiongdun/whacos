package m_file

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

type File struct {
	models.Model
	Name    string `json:"name"`    // 文件名
	KeyId   string `json:"keyId"`   // keyId
	Suffix  string `json:"suffix"`  // 文件后缀
	Size    int    `json:"size"`    // 文件大小
	Address string `json:"address"` // 文件地址
}

// 查询指定文件记录
func (f *File) SelectById(id int) (*File, error) {
	var file File
	if err := models.DB.Model(&file).Where("id = ? and del_flag = 1", id).Find(&file).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &file, nil
}

// 分页查询列表记录
func (f *File) SelectPage(param File, pageNum int, pageSize int) ([]File, error) {
	var files []File
	if err := models.DB.Where(&param).Order("file.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&files).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return files, nil
}

// 查询文件记录列表
func (f *File) SelectList(param File) ([]File, error) {
	var files []File
	if err := models.DB.Where(&param).Find(&files).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return files, nil
}

// 统计文件记录
func (f *File) Count(param File) (int, error) {
	var count int
	if err := models.DB.Model(&File{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除文件记录
func (f *File) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增文件记录
func (f *File) Insert(file File) error {
	if err := models.DB.Create(&file).Error; err != nil {
		return err
	}
	return nil
}

// 修改文件记录
func (f *File) UpdateById(file File) error {
	if err := models.DB.Model(&File{}).Where("id = ?", file.Id).Update(&file).Error; err != nil {
		return err
	}
	return nil
}
