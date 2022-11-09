package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	args := os.Args
	if args == nil {
		log.Fatal("no args provided")
	}
	if len(args) == 1 {
		fmt.Println("please provide arguments. (e.g. help)")
		return
	}
	fmt.Println(args)

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	
	r := gin.Default()

	r.GET("/", index)

	r.Run()
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello index",
	})
}
