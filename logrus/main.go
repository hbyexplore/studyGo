package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

/*
在 gin中结合 logrus使用
*/
//创建一个全局日志变量
var log = logrus.New()

func main() {
	//初始化logrus的配置
	//以json格式输出
	log.Formatter = &logrus.JSONFormatter{}
	//指定日志输出文件
	file, err := os.OpenFile("./logrus.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open file failed,err:", err)
	}
	//将日志的输出指定到文件
	log.Out = file
	//将gin框架的日志也记录到指定的日志文件中
	gin.SetMode(gin.ReleaseMode) //将gin的模式设置为线上
	gin.DefaultWriter = log.Out
	gin.DefaultErrorWriter = log.Out
	//设置日志级别
	log.Level = logrus.DebugLevel
	router := gin.Default()
	router.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
		log.Info("hello请求执行完毕")
	})
	router.Run()
}

/*
简单尝试第三方日志包使用 logrus
*/
func helloworld() {
	//将调用方法作为日志字段
	logrus.SetReportCaller(true)
	//修改日志输出结构为 json结构 默认是 text
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//记录结构体的日志方便日志检索
	logrus.WithFields(logrus.Fields{
		"name": "哈哈哈",
		"age":  212,
	}).Info("插入数据失败数据信息为:")
}
