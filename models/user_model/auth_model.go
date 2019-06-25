package user_model

type Auth struct {
	Id int `json:"id" gorm:"primary_key"`

	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	//var user User
	//models.DB.Model(&user).Table("sys_user user").Where(User{Username: username, Password: password}).Find(&user)
	//if user.Id > 0 {
	//	return true
	//}

	return true
}
