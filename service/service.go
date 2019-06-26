package service

import "time"

type BaseDTO struct {
	Id          int
	CreatedBy   int
	CreatedTime time.Time
	UpdatedBy   int
	UpdatedTime time.Time
	DelFlag     int
}

type BaseService interface {
	// 按ID 查询
	// @Param id
	FindById(id int) interface{}

	FindList(params interface{}) []interface{}

	Save(param interface{}) bool

	Create(param interface{}) bool

	Modify(param interface{}) bool

	Remove(id int) bool
}
