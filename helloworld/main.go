package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
GIN 路由 用的是 httpRouter 组件  路由速度很快 因为 组件内用的是  url前缀数进行搜索
*/
func main() {
	//启动一个默认的路由
	router := gin.Default()
	//给url配置一个函数
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "Hello GIN"})
	})
	//启动服务器监听一个端口,默认8080
	router.Run()
}
