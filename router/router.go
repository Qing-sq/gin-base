package router

import (
	"awesomeProject3/controllers"
	"awesomeProject3/middleware"
	"awesomeProject3/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register", controllers.UserRegister)
		v1.POST("user/login", controllers.UserLogin)
		v1.GET("user/getUser", controllers.GetUserList)
	}
	{
		v1.GET("good",controllers.GetGood)
	}
	{
		v1.GET("order",controllers.Ping)
	}
	v1.Use(jwtAutoMiddleware)
	{
		v1.POST("upload", controllers.Upload)
		v1.PUT("user/change",controllers.UserChange)
	}
	{
		v1.POST("good", controllers.AddGood)
	}
	return r
}

func jwtAutoMiddleware(c *gin.Context) {
	autoHeader := c.Request.Header.Get("Authorization")
	if autoHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    403,
			"message": "解析token失败",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(autoHeader," ",2)
	if !(len(parts) == 2 && parts[0] == "Bearer"){
		c.JSON(http.StatusOK,gin.H{
			"code":403,
			"message":"token格式有误",
		})
		c.Abort()
		return
	}
	//parts[1] 我们使用之前定义好解析jwt的函数来解析他
	mc,err := jwt.ParseToken(parts[1])
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code":403,
			"message":"无效token",

		})
		c.Abort()
		return
	}
	c.Set("phone",mc.Phone)
	c.Set("name",mc.Name)
	c.Next()
}
