package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func (context *gin.Context) {
		fmt.Println("已经授权")
		cookie, err := context.Request.Cookie("user_cookie")
		if err == nil {
			context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
		context.Next()
	}
}
