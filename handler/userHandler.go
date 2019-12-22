package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户名: "+username+"已保存")
}

func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户: "+username+"年龄: "+age+"已保存")
}
