package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	InitRouter(r)
	r.Run()
}
