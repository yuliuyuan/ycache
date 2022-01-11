package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yly/ycache/handler"
)

func InitRouter(r *gin.Engine) {
	//test
	r.POST("/boe/app", app)
	
	r.POST("/set", handler.Set)
	r.GET("/get/:key", handler.Get)
	r.DELETE("/del/:key", handler.Del)
	r.GET("/getstat", handler.GetStat)
}

//boe小程序test
func app(c *gin.Context) {
	fmt.Println(c)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
