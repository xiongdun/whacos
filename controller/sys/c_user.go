package sys

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"whacos/models"
	"whacos/models/user_model"
	"whacos/pkg/err"
)

// @Summary 按ID 获取用户
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/user/get [GET]
func GetUser(c *gin.Context) {

	// 取值
	id, _ := com.StrTo(c.Param("id")).Int()

	// 校验
	valid := validation.Validation{}
	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	// 查询
	user := user_model.SelectUserById(id)

	// 返回
	c.JSON(http.StatusOK, gin.H{
		"code": err.Success,
		"msg":  "请求成功！",
		"data": user,
	})
}

func ListUser(c *gin.Context) {

	maps := make(map[string]interface{})
	//maps["status"] = 1
	userList := user_model.SelectUserList(1, 20, maps)

	c.JSON(http.StatusOK, gin.H{
		"code": err.Success,
		"msg":  "请求成功",
		"data": userList,
	})
}

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	email := c.Query("email")
	idCard := c.Query("idCard")
	sex, _ := com.StrTo(c.Query("sex")).Int()
	birth := c.Query("birth")

	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空！")
	valid.Required(email, "email").Message("电子邮箱不能位空！")
	valid.Required(idCard, "idCard").Message("身份证账号不能位空！")
	valid.Required(sex, "sex").Message("性别不能位空！")
	valid.Required(birth, "birth").Message("出身日期不能位空！")

	user := &user_model.User{
		Username: username,
		Email:    email,
		IdCard:   idCard,
		Sex:      sex,
		Birth:    time.Now(),
		Model: models.Model{
			CreatedBy:   1,
			CreatedTime: time.Now(),
			UpdatedBy:   1,
			UpdatedTime: time.Now(),
			DelFlag:     1,
		},
	}

	result := user_model.InsertUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code": err.Success,
		"msg":  "请求成功",
		"data": result,
	})
}

func EditUser(c *gin.Context) {

}

// @Summary 移除指定用户
// @Produce json
// @Param query username body string true "username"
// @Param query password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /remove [DELETE]
func RemoveUser(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()

	valid := validation.Validation{}

	valid.Required(id, "id").Message("Id不能为空！")
	valid.Min(id, 1, "id").Message("Id必须大于0！")

	isSuccess := user_model.DeleteUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"code": err.Success,
		"msg":  "请求成功",
		"data": isSuccess,
	})
}
