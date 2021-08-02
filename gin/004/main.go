package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	// panic だとよ
	r.GET("/pong", func(c *gin.Context) {
		c.Header("Content-Type", "application/html; charset=utf-8")
		c.String(http.StatusOK, "Today is Ok")
	})

	r.GET("/pong", func(c *gin.Context) {
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.String(http.StatusOK, "Today is Ok")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
