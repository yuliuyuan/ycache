package handler

import (
	"github.com/gin-gonic/gin"
	"yly/ycache/cache"
)

type HttpHandler struct {
	cache cache.Cache
}

var (
	h *HttpHandler
)

func init() {
	h = &HttpHandler{
		cache: cache.New("inmemory"),
	}
}

func Set(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	h.cache.Set(key, []byte(value))
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func Get(c *gin.Context) {
	key := c.Param("key")
	value, err := h.cache.Get(key)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, gin.H{
		"value": string(value),
	})
}

func Del(c *gin.Context) {
	key := c.Param("key")
	err := h.cache.Del(key)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetStat(c *gin.Context) {
	stat := h.cache.GetStat()
	c.JSON(200, gin.H{
		"message": stat,
	})
}
