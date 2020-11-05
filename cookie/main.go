package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Any("/login")
	router.GET("/login")
	router.GET("/login")
	router.Run()
}
