package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 提交的文章内容
// @Id 1
// @Tags 文章
// @Version 1.0
// @Accept application/x-json-stream
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回
// @Failure 409 object model.Result 添加失败
// @Router /article [post]
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
