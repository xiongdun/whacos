package m_dict

import (
	"github.com/jinzhu/gorm"
	"whacos/models"
)

type Dict struct {
	models.Model

	Name     string `json:"name"`
	Value    string `json:"value"`
	DictType string `json:"dictType"`
	SortNum  int    `json:"sortNum"`
	ParentId int    `json:"parentId" gorm:"idx_parent_id"`
}

// 查询指定数据字典记录
func (d *Dict) SelectById(id int) (*Dict, error) {
	var dict Dict
	if err := models.DB.Model(&dict).Where("id = ? and del_flag = 1", id).Find(&dict).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &dict, nil
}

// 分页查询列表记录
func (d *Dict) SelectPage(param Dict, pageNum int, pageSize int) ([]Dict, error) {
	var dicts []Dict
	if err := models.DB.Where(&param).Order("dict.created_time desc").Offset(pageNum - 1).Limit(pageSize).Find(&dicts).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dicts, nil
}

// 查询数据字典记录列表
func (d *Dict) SelectList(param Dict) ([]Dict, error) {
	var dicts []Dict
	if err := models.DB.Where(&param).Find(&dicts).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dicts, nil
}

// 统计数据字典记录
func (d *Dict) Count(param Dict) (int, error) {
	var count int
	if err := models.DB.Model(&Dict{}).Where(&param).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 删除数据字典记录
func (d *Dict) DeleteById(id int) error {
	if err := models.DB.Where("id = ?", id).Update("del_flag", models.DelFlagNo).Error; err != nil {
		return err
	}
	return nil
}

// 新增数据字典菜单记录
func (d *Dict) Insert(dict Dict) error {
	if err := models.DB.Create(&dict).Error; err != nil {
		return err
	}
	return nil
}

// 修改数据字典记录
func (d *Dict) UpdateById(dict Dict) error {
	if err := models.DB.Model(&Dict{}).Where("id = ?", dict.Id).Update(&dict).Error; err != nil {
		return err
	}
	return nil
}
