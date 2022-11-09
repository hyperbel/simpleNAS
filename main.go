package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	handleArgs(os.Args)

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

func handleArgs(args []string) string {
	
	fmt.Println(args)
	fmt.Println(len(args))
	if args == nil {
		home_dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("user doesn't have home dir...")
			os.Exit(1)
		}
		fmt.Println(home_dir)
		//config_file_location := fmt.Sprintf("%d/.config/simplenas/config", home_dir)
	}

	if len(args) == 1 {
		fmt.Println("please provide arguments. (e.g. help)")
		os.Exit(1)
	}

	if args[1] == "help" {
		fmt.Println("You can provide a config file, by passing it as an Argument.")
		fmt.Println("Usage: go run . <config>")
		fmt.Println("the default config file is ~/.config/simplenas/config")
		fmt.Println("If this config file is not there, you have to provide one")
	}
	
//	return config_file_location
	return ""
}

