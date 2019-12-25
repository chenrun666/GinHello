package initRouter

import (
	"GinHello/handler"
	"GinHello/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 加载模板
	router.LoadHTMLGlob("templates/*")
	// 添加图标
	router.StaticFile("/favicon.ico", "./favicon.ico")
	// 加载静态资源
	router.Static("/statics", "./statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath() + "avatar/"))

	// 添加路由
	index := router.Group("/")
	{
		index.Any("", retHelloGinAndMethod)
	}

	// User 路由
	userRouter := router.Group("/user")
	{
		userRouter.GET("", handler.UserSaveByQuery)
		//userRouter.GET("/:name", handler.UserSave)
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile", handler.UserProfile)
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	// Index
	indexRouter := router.Group("/index")
	{
		indexRouter.GET("", handler.IndexHandler)
	}

	return router
}
