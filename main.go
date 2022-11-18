package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var Conf Config

func main() {
	file_path := handleArgs(os.Args)
	json_file, err := os.Open(file_path)

	if err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	byte_value, _ := io.ReadAll(json_file)
	json.Unmarshal(byte_value, &Conf)

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	
	r.LoadHTMLGlob("sites/html/*.html")
	r.Static("/assets", "./sites/assets")

	r.GET("/", index)
	r.GET("/dir", dir)
	r.POST("/login", login)
	r.POST("/createaccount", createaccount)
	r.POST("/back", back)

	r.Run()
}

