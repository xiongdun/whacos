package c_user

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"time"
	"whacos/models"
	"whacos/models/m_user"
	"whacos/pkg/app"
	"whacos/pkg/logging"
	"whacos/pkg/utils"
	"whacos/service/s_user"
)

type ListUserForm struct {
	app.CommonForm
	Username string `form:"username" json:"username"`
	Name     string `form:"name" json:"name"`
	IdCard   string `form:"idCard" json:"idCard" valid:"MaxSize(18)"`
}

type AddUserForm struct {
	Username string `form:"username" json:"username" valid:"Required;MaxSize(50)"`
	Name     string `form:"name" json:"name" valid:"Required;MaxSize(18)"`
	IdCard   string `form:"idCard" json:"idCard" valid:"Required;MaxSize(18)"`
	Email    string `form:"email" json:"email" valid:"Required;MaxSize(64)"`
	Mobile   string `form:"mobile" json:"mobile" valid:"Required;MaxSize(15)"`
	Sex      int    `form:"sex" json:"sex" valid:"Required;Range(0,1)"`
}

type EditUserForm struct {
	Id          int    `json:"id" valid:"Required;MinSize(1)"`
	Username    string `swaggo:"true,用户名" json:"username" valid:"Required;MaxSize(50)"`
	Name        string `json:"name" valid:"Required;MaxSize(50)"`
	DeptId      int    `json:"deptId" valid:"Required;MaxSize(20)"`
	Email       string `json:"email" valid:"Required;MaxSize(64)"`
	Mobile      string `json:"mobile" valid:"Required;MaxSize(15)"`
	IdCard      string `json:"idCard" valid:"Required;MaxSize(18)"`
	Status      int    `json:"status" valid:"Required;Range(0,1,2)"`
	Sex         int    `json:"sex" valid:"Required;Range(0,1)"`
	PicId       int    `json:"picId" valid:"Required;MaxSize(20)"`
	LiveAddress string `json:"liveAddress" valid:"Required;MaxSize(50)"`
	Hobby       string `json:"hobby" valid:"Required;MaxSize(50)"`
	Province    string `json:"province" valid:"Required;MaxSize(50)"`
	City        string `json:"city" valid:"Required;MaxSize(50)"`
	District    string `json:"district" valid:"Required;MaxSize(50)"`
	Remarks     string `json:"remarks" valid:"Required;MaxSize(255)"`
}

// @Tags 系统用户相关
// @Resource Name
// @Summary 按ID 获取用户
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/user/get/{id} [GET]
func GetUser(c *gin.Context) {
	appGin := app.Gin{C: c}
	// 取值
	id, _ := com.StrTo(c.Param("id")).Int()

	// 校验
	valid := validation.Validation{}
	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	if valid.HasErrors() {
		appGin.ResponseInvalidParams(nil)
		return
	}

	reply, err := utils.SetString("xing", "sdsdsd")
	if err == nil {
		fmt.Println(reply)
	}
	value, err1 := utils.GetString("xing")
	if err1 == nil {
		fmt.Println(value)
	}

	// 查询并返回

	if user, err := s_user.FindUserById(id); err != nil {
		appGin.ResponseFail(err.Error())
	} else {
		appGin.ResponseSuccess(user)
	}

}

// @Tags APP相关
// @Resource Name
// @Summary 按ID 获取用户
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param token query string true "token"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /app/user/get/{id} [GET]
func GetAppUser(c *gin.Context) {
	appGin := app.Gin{C: c}
	// 取值
	id, _ := com.StrTo(c.Param("id")).Int()

	// 校验
	valid := validation.Validation{}
	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	if valid.HasErrors() {
		appGin.ResponseInvalidParams(nil)
		return
	}

	// 查询并返回

	if user, err := s_user.FindUserById(id); err != nil {
		appGin.ResponseFail(err.Error())
	} else {
		appGin.ResponseSuccess(user)
	}

}

// @Tags 系统用户相关
// @Resource Name
// @Summary 查询用户列表
// @Accept json
// @Produce json
// @Param listUser body c_user.ListUserForm true "listUserForm"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/user/list [POST]
func ListUser(c *gin.Context) {

	appGin := app.Gin{C: c}
	userForm := ListUserForm{}
	if c.BindJSON(&userForm) != nil {
		appGin.ResponseInvalidParams(nil)
		return
	}

	user := m_user.SysUser{
		Username: userForm.Username,
		Name:     userForm.Name,
	}

	if userPage, err := s_user.FindUserPage(user, userForm.PageNum, userForm.PageSize); err != nil {
		appGin.ResponseFail(nil)
	} else {
		appGin.ResponseSuccess(userPage)
	}
}

// @Tags 系统用户相关
// @Resource Name
// @Summary 新增用户
// @Accept json
// @Produce json
// @Param addUser body c_user.AddUserForm true "addUserForm"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/user/add [POST]
func AddUser(c *gin.Context) {
	appGin := app.Gin{C: c}

	addUserForm := AddUserForm{}

	if err := c.BindJSON(&addUserForm); err != nil {
		appGin.ResponseInvalidParams(nil)
		logging.Info(err.Error())
		return
	}

	user := m_user.SysUser{
		Username: addUserForm.Username,
		Name:     addUserForm.Name,
		IdCard:   addUserForm.IdCard,
		Mobile:   addUserForm.Mobile,
		Sex:      addUserForm.Sex,
		Email:    addUserForm.Email,
		Status:   1,
		Birth:    time.Now(),
		Model: models.Model{
			DelFlag:     models.DelFlagYes,
			CreatedBy:   1,
			CreatedTime: time.Now(),
			UpdatedBy:   1,
			UpdatedTime: time.Now(),
		},
	}

	if err := s_user.CreateUser(user); err != nil {
		appGin.ResponseFail(nil)
	} else {
		appGin.ResponseSuccess(true)
	}

}

// @Tags 系统用户相关
// @Resource Name
// @Summary 修改用户信息
// @Accept json
// @Produce json
// @Param editUser body c_user.EditUserForm true "editUserForm"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/user/edit [PUT]
func EditUser(c *gin.Context) {
	appGin := app.Gin{C: c}

	editUserForm := EditUserForm{}

	if err := c.BindJSON(&editUserForm); err != nil {
		appGin.ResponseInvalidParams(nil)
		logging.Info(err.Error())
		return
	}

	user := m_user.SysUser{
		Username:    editUserForm.Username,
		Name:        editUserForm.Name,
		IdCard:      editUserForm.IdCard,
		Mobile:      editUserForm.Mobile,
		Sex:         editUserForm.Sex,
		Email:       editUserForm.Email,
		Status:      editUserForm.Status,
		LiveAddress: editUserForm.LiveAddress,
		Hobby:       editUserForm.Hobby,
		Province:    editUserForm.Province,
		City:        editUserForm.City,
		District:    editUserForm.District,
		Remarks:     editUserForm.Remarks,
		//Birth:    none,
		Model: models.Model{
			//DelFlag:     1,
			//CreatedBy:   1,
			//CreatedTime: time.Now(),
			UpdatedBy:   1,
			UpdatedTime: time.Now(),
			Id:          editUserForm.Id,
		},
	}

	if err := s_user.ModifyUser(user); err != nil {
		appGin.ResponseFail(nil)
	} else {
		appGin.ResponseSuccess(true)
	}
}

// @Tags 系统用户相关
// @Resource Name
// @Summary 移除指定用户
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/user/remove/{id} [DELETE]
func RemoveUser(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()

	valid := validation.Validation{}

	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	appGin := app.Gin{C: c}
	if valid.HasErrors() {
		appGin.ResponseInvalidParams(nil)
		logging.Info(valid.Errors[0].Message)
		return
	}

	if err := s_user.RemoveUserById(id); err != nil {
		appGin.ResponseFail(err.Error())
	} else {
		appGin.ResponseSuccess(true)
	}
}
