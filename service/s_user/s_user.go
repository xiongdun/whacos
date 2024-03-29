package s_user

import (
	"fmt"
	"whacos/models"
	"whacos/models/m_user"
	"whacos/pkg/utils"
)

// 按用户ID查询
// param id 用户ID
func FindUserById(id int) (*m_user.SysUser, error) {
	var sysUser m_user.SysUser
	return sysUser.SelectById(id)
}

// 查询用户列表
func FindUserList(param m_user.SysUser) ([]m_user.SysUser, error) {
	return param.SelectList(param)
}

// 分页查询
func FindUserPage(param m_user.SysUser, pageNum int, pageSize int) (p models.Page, err error) {
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
func SaveUser(param m_user.SysUser) error {
	if param.Id != 0 {
		return param.UpdateById(param)
	} else {
		return param.Insert(param)
	}
}

// 统计用户数
func CountUser(param m_user.SysUser) (int, error) {
	if count, err := param.Count(param); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

// 创建用户
func CreateUser(param m_user.SysUser) error {
	return param.Insert(param)
}

// 修改用户
func ModifyUser(param m_user.SysUser) error {
	return param.UpdateById(param)
}

// 按用户ID删除
// param id 用户ID
func RemoveUserById(id int) error {
	sysUser := m_user.SysUser{}
	return sysUser.DeleteById(id)
}

// 查询校验
func CheckAuth(username, password string) bool {
	sysUser := m_user.SysUser{}
	if result, err := sysUser.SelectByUsername(username); err != nil {
		return false
	} else {
		// md5 加密
		md5 := utils.EncodeMD5(password)
		fmt.Println(md5)
		if result != nil && result.Password == md5 {
			return true
		}
	}
	return false
}
