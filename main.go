package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Dir	string	`json:"dir"`
}

func main() {
	file_path := handleArgs(os.Args)
	json_file, err := os.Open(file_path)
	if err !=nil {
		log.Fatal(err)
		os.Exit(1)
	}
	byte_value, _ := ioutil.ReadAll(json_file)
	var conf Config
	json.Unmarshal(byte_value, &conf)
	fmt.Println(conf)

	r := gin.Default()
	
	r.LoadHTMLGlob("sites/html/*.html")

	r.GET("/", index)
	r.GET("/dirtest", func(c *gin.Context) {
		r.LoadHTMLFiles("sites/html/dirtest.html")
		c.HTML(http.StatusOK, "dirtest.html", gin.H{
			"message": "directory test stuff",
		})
	})

	r.Run()
}
func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "index called",
	})
}


func handleArgs(args []string) string {
	config_file_location := ""
	home_dir, err := os.UserHomeDir()

	if len(args) == 1 {
		if err != nil {
			fmt.Println("user doesn't have home dir...")
			os.Exit(1)
		}
		config_file_location = fmt.Sprintf("%s/.config/simplenas/config.json", home_dir)
		fmt.Println(config_file_location)
		return config_file_location
	}

	if args[1] == "help" {
		fmt.Println("You can provide a config file, by passing it as an Argument.")
		fmt.Println("Usage: go run . <config>")
		fmt.Println("the default config file is ~/.config/simplenas/config")
		fmt.Println("If this config file is not there, you have to provide one")
		os.Exit(0)
	} else {
		config_file_location = args[1]
		fmt.Println(args[1])
	}
	
	return config_file_location
}
