package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	/*
		获取 get请求的参数
	*/
	router.GET("/query", func(context *gin.Context) {
		//获取query中的 name参数
		//queryName :=context.Query("name")
		queryName, _ := context.GetQuery("name")
		//设置默认值,没取到对应的值就取默认值
		defaultAge := context.DefaultQuery("age", "默认值")
		context.JSON(http.StatusOK, gin.H{
			"name": queryName,
			"age":  defaultAge,
		})
	})

	/*
		获取 post请求的值
	*/
	router.POST("/form", func(context *gin.Context) {
		nameVal := context.PostForm("name")
		age := context.DefaultPostForm("age", "19")
		context.JSON(http.StatusOK, gin.H{
			"name": nameVal,
			"age":  age,
		})
	})

	/*
		获取get请求url中的值
	*/
	router.GET("/book/:param", func(context *gin.Context) {
		paramVal := context.Param("param")
		context.JSON(http.StatusOK, gin.H{
			"param": paramVal,
		})
	})
	router.Run()
}
