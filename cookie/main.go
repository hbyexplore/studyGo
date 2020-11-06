package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

var userInfo user

func main() {
	router := gin.Default()
	//加载html文件 全局的
	router.LoadHTMLGlob("./html/*")
	router.Use(middlewareCookie)
	router.Any("/login", loginHandler)
	router.GET("/index", indexHandler)
	router.GET("/cart", cartHandler)
	router.GET("/userInfo", userInfoHandler)
	router.Run()
}
func middlewareCookie(context *gin.Context) {
	url := context.Request.URL.Path
	//如果是去登录就放行
	if url == "/login" {
		context.Next()
	} else {
		//查看cookie中有没有值
		value, err := context.Cookie("name")
		if err != nil || value == "" {
			//重定向状态码必须是 301
			context.Redirect(http.StatusMovedPermanently, "/login")
			return
		}
		context.Next()
	}

}
func loginHandler(context *gin.Context) {
	if context.Request.Method == "POST" {
		err := context.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println("绑定参数出错,", err)
		}
		if userInfo.Name == "" && userInfo.Password == "" {
			context.HTML(http.StatusOK, "login.html", gin.H{
				"msg": "用户名或密码不能为空",
			})
		}
		//用户名和密码正确 生成cookie  取cookie的格式就是  domain+path+当前的url 例如:localhost:8080/index 如果path和domain不对的话无法找到cookie
		context.SetCookie("name", userInfo.Name, 60, "/", "127.0.0.1", false, true)
		//重定向状态码必须是 301
		context.Redirect(http.StatusMovedPermanently, "/index")
	} else {
		context.HTML(http.StatusOK, "login.html", nil)
	}
}
func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
func cartHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "cart.html", gin.H{
		"name": userInfo.Name,
	})
}
func userInfoHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "userInfo.html", gin.H{
		"name": userInfo.Name,
	})
}
