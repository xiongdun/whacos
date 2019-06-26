package c_user

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/models/m_user"
	"whacos/pkg/app"
	"whacos/pkg/e"
	"whacos/pkg/logging"
	"whacos/service/s_user"
)

type ListUserForm struct {
	app.CommonForm
	Username string `form:"username" json:"username"`
	Name     string `form:"name" json:"name"`
	IdCard   string `form:"idCard" json:"idCard" valid:"MaxSize(18)"`
}

type AddUserForm struct {
	Username string `form:"username" json:"username" valid:"Required MaxSize(50)"`
	Name     string `form:"name" json:"name" valid:"Required MaxSize(18)"`
	IdCard   string `form:"idCard" json:"idCard" valid:"Required MaxSize(18)"`
	Email    string `form:"email" json:"email" valid:"Required MaxSize(64)"`
	Mobile   string `form:"mobile" json:"mobile" valid:"Required MaxSize(15)"`
	Sex      string `form:"sex" json:"sex" valid:"Required MaxSize(1)"`
}

type EditUserForm struct {
	Username    string `swaggo:"true,用户名" json:"username" valid:"Required MaxSize(50)"`
	Name        string `json:"name" valid:"Required MaxSize(50)"`
	DeptId      int    `json:"deptId" valid:"Required MaxSize(20)"`
	Email       string `json:"email" valid:"Required MaxSize(64)"`
	Mobile      string `json:"mobile" valid:"Required MaxSize(15)"`
	IdCard      string `json:"idCard" valid:"Required MaxSize(18)"`
	Status      int    `json:"status" valid:"Required MaxSize(1)"`
	Sex         int    `json:"sex" valid:"Required MaxSize(1)"`
	PicId       int    `json:"picId" valid:"Required MaxSize(20)"`
	LiveAddress string `json:"liveAddress" valid:"Required MaxSize(50)"`
	Hobby       string `json:"hobby" valid:"Required MaxSize(50)"`
	Province    string `json:"province" valid:"Required MaxSize(50)"`
	City        string `json:"city" valid:"Required MaxSize(50)"`
	District    string `json:"district" valid:"Required MaxSize(50)"`
	Remarks     string `json:"remarks" valid:"Required MaxSize(255)"`
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

	// 取值
	id, _ := com.StrTo(c.Param("id")).Int()

	// 校验
	valid := validation.Validation{}
	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	// 查询
	user := m_user.SelectUserById(id)

	// 返回
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功！",
		"data": user,
	})
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
	}

	maps := make(map[string]interface{})
	//maps["status"] = 1
	userList := m_user.SelectUserList(userForm.PageNum, userForm.PageSize, maps)

	appGin.ResponseSuccess(userList)
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
	user := s_user.UserDTO{}

	if result, err := user.Create(user); err != nil {
		appGin.ResponseFail(nil)
	} else {
		appGin.ResponseSuccess(result)
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
	id, _ := com.StrTo(c.Query("id")).Int()

	valid := validation.Validation{}

	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	isSuccess := m_user.DeleteUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
		"data": isSuccess,
	})
}
