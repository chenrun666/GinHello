package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
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

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err -> ", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		id := user.Save()
		log.Println("id is ", id)
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}
}
