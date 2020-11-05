package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
文件上传
*/
func main() {
	//创建一个默认路由
	router := gin.Default()
	//加载html文件
	router.LoadHTMLFiles("./upload.html")
	//创建跳转到指定html文件的get方法
	router.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "upload.html", nil)
	})
	//接收上传的文件的post请求
	router.POST("/submitFile", func(context *gin.Context) {
		//拿到file指针对象
		fileHeader, err := context.FormFile("file")
		//拼接文件保存路径
		dst := fmt.Sprintf("./%s", fileHeader.Filename)
		//保存文件
		err = context.SaveUploadedFile(fileHeader, dst)
		if err != nil {
			context.JSON(http.StatusInternalServerError, "上传文件失败")
		}
		//返回消息
		context.JSON(http.StatusOK, "上传文件成功")
	})

	//接收上传的多个文件的post请求
	router.POST("/multipartFile", func(context *gin.Context) {
		//拿到form指针对象
		form, _ := context.MultipartForm()
		//拿到所有的file指针对象数组
		for _, f := range form.File {
			//循环拿到所有的文件
			for index, fileHeader := range f {
				//拼接文件保存路径
				dst := fmt.Sprintf("./%s%d", fileHeader.Filename, index)
				//保存文件
				err := context.SaveUploadedFile(fileHeader, dst)
				if err != nil {
					context.JSON(http.StatusInternalServerError, "上传文件失败")
				}
			}
		}
		//返回消息
		context.JSON(http.StatusOK, "上传文件成功")
	})

	//开启服务器
	router.Run()
}
