package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
路由分组
*/
func main() {
	router := gin.Default()
	routerGroup := router.Group("/user")
	routerGroup.Use(interceptorOne)
	{
		routerGroup.GET("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/user/login",
			})
		})
		routerGroup.GET("/show")
		routerGroup.POST("/save")
	}
	routerGroup2 := router.Group("/image")
	routerGroup2.Use(interceptorTwo)
	{
		routerGroup2.GET("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/image/login",
			})
		})
		routerGroup2.GET("/show")
		routerGroup2.POST("/save")
	}

	//还可以多组嵌套
	v1 := router.Group("/v1")
	{
		v2 := v1.Group("v2")
		{
			v2.GET("login", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"msg": "/image/login",
				})
			})
		}
	}
	router.Run()
}
func interceptorOne(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Second * 5)
	//拦截器中必须执行 next() 方法否则无法继续走下面的执行器
	c.Next()
	fmt.Println("执行时间为:", time.Since(start))
}
func interceptorTwo(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Second * 3)
	//拦截器中必须执行 next() 方法否则无法继续走下面的执行器
	c.Next()
	fmt.Println("执行时间为:", time.Since(start))
}
