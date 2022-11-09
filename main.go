package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", index)

	r.Run()
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello index",
	})
}
