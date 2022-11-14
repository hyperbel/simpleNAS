package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
//	"encoding/json"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	files, err := os.ReadDir(Conf.Dir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	 fs := make([]FileInfo, len(files))	//change 100 to amount of files

	
	for i, f := range files {
		fs[i] = FileInfo{f.Name(), f.IsDir(), 0}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"dir": Conf.Dir,
		"files": fs,
	})

}

func dir(c *gin.Context) {
	path := c.Query("path")

	dir := Conf.Dir + path
	files, err := os.ReadDir(dir)
	
	if err != nil {
		log.Fatal(err)
		c.HTML(500, "error.html", gin.H{
			"message": "an error occured, please check logs",
		})
	}
	
	fs := make([]FileInfo, len(files))
	
	for i, file := range files {
		fs[i] = FileInfo{file.Name(), file.IsDir(), 0}
	}
	
	c.HTML(http.StatusOK, "dir.html", gin.H{
		"dir": dir,
		"files": fs,
	})
}

func login(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	username := c.PostForm("uname")
	password := c.PostForm("passwd")
//	user := User{0,username, password}
	fmt.Println(username, password)

	rows, err := db.Query("select name from Users where id=1;")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	match := false
	
	for rows.Next() {
		var p []byte
		err = rows.Scan(&p)

		if err != nil {
			log.Fatal(err)
		}
		/*
		var user User
		err := json.Unmarshal(p, &user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		*/
	}

	if match {
		c.Redirect(http.StatusMovedPermanently, "/dir?path=/")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func signin(c *gin.Context) {
	
}
