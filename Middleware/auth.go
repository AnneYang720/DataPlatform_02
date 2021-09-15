package middleware

import (
	"fmt"
	"net/http"

	"github.com/AnneYang720/DataBase2/App"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		G_userId := session.Get("userid")

		if G_userId != nil {
			fmt.Printf("authorize")
			c.Set("G_userId", G_userId)
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, &App.ApiResponse{
				Code:    400,
				Flag:    false,
				Message: "please login first",
			})
			return
		}
	}
}

// 中间件,主要处理js访问时的跨域问题
func Passjs() gin.HandlerFunc {
	return func(c *gin.Context) {

		// gin设置响应头，设置跨域

		c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Credentials","true")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			// session := sessions.Default(c)
			// session.Set("sessionid", 123456)       //change
			// session.Save()                             //change
			//session := sessions.Default(c)
			//v := session.Get("sessionid")
			//fmt.Println(v)
			c.JSON(200, &App.ApiResponse{
				Code: 200,
			})
			return
		}
		// c.Next()后就执行真实的路由函数
		c.Next()
	}
}
