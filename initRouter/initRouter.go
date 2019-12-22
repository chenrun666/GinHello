package initRouter

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 添加路由
	index := router.Group("/")
	{
		index.Any("", retHelloGinAndMethod)
	}

	// User 路由
	userRouter := router.Group("/user")
	{
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.GET("/:name", handler.UserSave)
	}

	return router
}
