package user

import (
	"GinHello/config"
	"GinHello/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateJwt(context *gin.Context) {
	// 获取用户
	user := &model.UserModel{}
	result := &model.Result{
		Code:    200,
		Message: "登录成功",
		Data:    nil,
	}
	if e := context.BindJSON(&user); e != nil {
		result.Message = "数据绑定失败"
		result.Code = http.StatusUnauthorized
		context.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		expiresTime := time.Now().Unix() + int64(config.OneDayOfHous)
		calims := jwt.StandardClaims{
			Audience:  user.Email,        // 受众人
			ExpiresAt: expiresTime,       // 失效时间
			Id:        string(user.Id),   // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "gin hello",       // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}
		var jwtSecret = []byte(config.Secret)
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, calims)
		if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
			result.Message = "登录成功"
			result.Data = "Bearer " + token
			result.Code = http.StatusOK
			context.JSON(result.Code, gin.H{
				"result": result,
			})
		} else {
			result.Message = "登录失败"
			result.Code = http.StatusOK
			context.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "登录失败"
		result.Code = http.StatusOK
		context.JSON(result.Code, gin.H{
			"result": result,
		})
	}
}
