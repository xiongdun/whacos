package s_user

import (
	"time"
	"whacos/service"
)

type UserDTO struct {
	service.BaseDTO

	Username    string
	Name        string
	Password    string
	DeptId      int
	Email       string
	Mobile      string
	IdCard      string
	Status      int
	Sex         int
	Birth       time.Time
	PicId       int
	LiveAddress string
	Hobby       string
	Province    string
	City        string
	District    string
	Remarks     string
}

func (u *UserDTO) FindById(id int) (user UserDTO) {
	return
}

func (u *UserDTO) FindList(param UserDTO) (users []UserDTO) {
	return
}

func (u *UserDTO) Save(param UserDTO) (b bool) {
	return
}

func (u *UserDTO) Count(param UserDTO) (count int) {
	return
}
func (u *UserDTO) Create(param UserDTO) (bool, error) {
	var err error
	return true, err
}

func (u *UserDTO) Modify(param UserDTO) (b bool) {
	return
}

func (u *UserDTO) Remove(id int) (b bool) {
	return
}
