package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"whacos/controller/api"
	"whacos/controller/sys"
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
	router.GET("/", sys.Index)
	// jwt token 校验
	router.POST("/auth", api.GetAuth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/login", sys.ToLogin)
	router.POST("/login", sys.Login)

	appApi := router.Group("/api")
	appApi.Use(jwt.ValidJWT())
	{
		appApi.GET("/file/upload", api.UploadFile)
	}

	sysUser := router.Group("/sys/user")
	{
		sysUser.GET("/get/:id", sys.GetUser)

		sysUser.POST("/list", sys.ListUser)

		sysUser.POST("/add", sys.AddUser)

		sysUser.PUT("/edit/:id", sys.EditUser)

		sysUser.DELETE("/remove/:id", sys.RemoveUser)
	}
	// 角色
	sysRole := router.Group("/sys/role")
	{
		sysRole.GET("/get", sys.GetRole)

		sysRole.POST("/list", sys.ListRole)

		sysRole.POST("/add", sys.AddRole)

		sysRole.PUT("/edit", sys.EditRole)

		sysRole.DELETE("/remove", sys.RemoveRole)
	}
	// 菜单
	sysMenu := router.Group("/sys/menu")
	{
		sysMenu.GET("/get", sys.GetMenu)

		sysMenu.POST("/list", sys.ListMenu)

		sysMenu.POST("/add", sys.AddMenu)

		sysMenu.PUT("/edit", sys.EditMenu)

		sysMenu.DELETE("/remove", sys.RemoveMenu)
	}

	return router
}
