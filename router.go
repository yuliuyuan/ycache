package main

import (
	"github.com/gin-gonic/gin"
	"yly/ycache/handler"
)

func InitRouter(r *gin.Engine) {
	r.POST("/set", handler.Set)
	r.GET("/get/:key", handler.Get)
	r.DELETE("/del/:key", handler.Del)
	r.GET("/getstat", handler.GetStat)

}

