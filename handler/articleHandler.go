package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertArticle(context *gin.Context) {
	article := model.Article{}
	var id = -1

	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
