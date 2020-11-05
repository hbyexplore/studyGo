package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//字段名必须大写  因为是通过反射赋值
type student struct {
	Name string `json:"name" form:"name"` //binding:"required" 表示该字段请求中必须传 不传报错
	Age  int    `json:"age" form:"age"`
}

var students student

func main() {
	//获取一个默认的引擎实例
	router := gin.Default()
	//创建一个实体类
	//get请求数据映射实体类
	router.GET("bindGet", getHandler)
	//post 请求数据映射实体类
	router.POST("bindPost", postHandler)
	//路由重定向
	router.GET("redirectRouter", func(context *gin.Context) {
		//路由重定向
		context.Request.URL.Path = "/bindGet"
		router.HandleContext(context)
	})
	//启动服务
	router.Run()
}
func getHandler(context *gin.Context) {
	if err := context.ShouldBind(&students); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "error__students"})
	} else {
		//重定向
		context.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
		/*context.JSON(http.StatusOK, gin.H{
			"name": students.Name,
			"age":  students.Age,
		})*/
	}
}
func postHandler(context *gin.Context) {
	if err := context.ShouldBind(&students); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"name": students.Name,
			"age":  students.Age,
		})
	}
}
