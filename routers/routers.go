package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"whacos/controller/c_file"
	"whacos/controller/c_index"
	"whacos/controller/c_menu"
	"whacos/controller/c_role"
	"whacos/controller/c_user"
	_ "whacos/docs"
	"whacos/middleware/jwt"
	"whacos/pkg/settings"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// 设置运行环境
	gin.SetMode(settings.RunMode)

	// 设置页面地址
	router.LoadHTMLGlob("templates/*")
	// 首页
	router.GET("/", c_index.Index)
	// jwt token 校验
	router.POST("/auth", c_user.GetAuth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/login", c_index.ToLogin)
	router.POST("/login", c_index.Login)

	appApi := router.Group("/app")
	appApi.Use(jwt.ValidJWT())
	{
		appApi.GET("/file/upload", c_file.UploadFile)
	}

	sysUser := router.Group("/sys/user")
	{
		sysUser.GET("/get/:id", c_user.GetUser)

		sysUser.POST("/list", c_user.ListUser)

		sysUser.POST("/add", c_user.AddUser)

		sysUser.PUT("/edit", c_user.EditUser)

		sysUser.DELETE("/remove/:id", c_user.RemoveUser)
	}
	// 角色
	sysRole := router.Group("/sys/role")
	{
		sysRole.GET("/get", c_role.GetRole)

		sysRole.POST("/list", c_role.ListRole)

		sysRole.POST("/add", c_role.AddRole)

		sysRole.PUT("/edit", c_role.EditRole)

		sysRole.DELETE("/remove", c_role.RemoveRole)
	}
	// 菜单
	sysMenu := router.Group("/sys/menu")
	{
		sysMenu.GET("/get", c_menu.GetMenu)

		sysMenu.POST("/list", c_menu.ListMenu)

		sysMenu.POST("/add", c_menu.AddMenu)

		sysMenu.PUT("/edit", c_menu.EditMenu)

		sysMenu.DELETE("/remove", c_menu.RemoveMenu)
	}

	return router
}
