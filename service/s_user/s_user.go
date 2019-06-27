package s_user

import (
	"whacos/models"
	"whacos/models/m_user"
)

// 按用户ID查询
// param id 用户ID
func FindUserById(id int) (*m_user.User, error) {
	var user m_user.User
	return user.SelectById(id)
}

// 查询用户列表
func FindUserList(param m_user.User) ([]m_user.User, error) {
	return param.SelectList(param)
}

// 分页查询
func FindUserPage(param m_user.User, pageNum int, pageSize int) (p models.Page, err error) {
	if userPage, err := param.SelectPage(param, pageNum, pageSize); err != nil {
		return models.PageUtils(nil, 0, pageSize), err
	} else {
		if count, err2 := CountUser(param); err2 == nil {
			return models.PageUtils(userPage, count, pageSize), nil
		}
	}
	return
}

// 保存用户
func SaveUser(param m_user.User) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计用户数
func CountUser(param m_user.User) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建用户
func CreateUser(param m_user.User) error {
	return param.Insert(param)
}

// 修改用户
func ModifyUser(param m_user.User) error {
	return param.UpdateById(param)
}

// 按用户ID删除
// param id 用户ID
func RemoveUserById(id int) error {
	user := m_user.User{}
	return user.DeleteById(id)
}

func CheckAuth(username, password string) bool {
	//var user User
	//models.DB.Model(&user).Table("sys_user user").Where(User{Username: username, Password: password}).Find(&user)
	//if user.Id > 0 {
	//	return true
	//}

	return true
}
