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
	fmt.Println("login called")
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("opening db successfull")

	rows, err := db.Query("select name from Users where id=1;")
	fmt.Println("got rows and err")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("gettings rows successful")
	match := false
	fmt.Println("match set to false")
	
	for rows.Next() {
		fmt.Println("looping through rows n")
		var p []byte
		err = rows.Scan(&p)
		fmt.Println("scanned row")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("scanning rows success")
		fmt.Println(p)
		/*
		var user User
		err := json.Unmarshal(p, &user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		*/
		fmt.Println(string(p[:]))
	}

	if match {
		c.Redirect(http.StatusMovedPermanently, "/dir?path=/")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func signin(c *gin.Context) {
	
}
