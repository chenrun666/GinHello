package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}
