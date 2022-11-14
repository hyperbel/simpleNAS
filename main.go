package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

var Conf Config

func main() {
	file_path := handleArgs(os.Args)
	json_file, err := os.Open(file_path)

	if err !=nil {
		log.Fatal(err)
		os.Exit(1)
	}

	byte_value, _ := ioutil.ReadAll(json_file)
	json.Unmarshal(byte_value, &Conf)


	r := gin.Default()
	
	r.LoadHTMLGlob("sites/html/*.html")

	r.GET("/", index)
	r.GET("/dir", dir)
	r.POST("/login", login)
	r.POST("/signin", signin)

	r.Run()
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
