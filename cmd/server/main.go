package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := setupRouter()
	err := r.Run(`:8080`)
	if err != nil {
		return
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
