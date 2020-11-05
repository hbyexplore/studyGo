package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
Gin的中间件(类似于拦截器)
*/
func main() {
	//拦截器可以使用 router设置n个
	router := gin.Default()
	router.Use(interceptorOne)
	//也可以直接在请求中设置
	//router.GET("/hello",interceptorOne,interceptorTwo)
	router.GET("/hello", interceptorTwo)
	router.GET("/hi", interceptorThree)
	router.Run()
}

//第一个请求执行方法
//计算请求执行时机
func interceptorOne(c *gin.Context) {
	start := time.Now()
	time.Sleep(time.Second * 5)
	//拦截器中必须执行 next() 方法否则无法继续走下面的执行器
	c.Next()
	fmt.Println("执行时间为:", time.Since(start))
}

func interceptorTwo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}
func interceptorThree(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hi",
	})
}
