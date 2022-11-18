package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"net/http"
	"crypto/sha256"
	"encoding/base64"
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
	db, err := sql.Open("sqlite3", Conf.DB)
	hasher := sha256.New()
	var match bool 
	var u User

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	username := c.PostForm("uname")
	password := c.PostForm("passwd")
	fmt.Println(c.Request.PostForm)
	fmt.Println(username, password)

	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		
	q, err := db.Prepare("SELECT * FROM Users WHERE name=? and password=?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)

	err = q.QueryRow(username, sha).Scan(&u.id, &u.name, &u.passwd)

	if err != nil {
	match = false		
	log.Println(err)
	} 	
	match = true
	
	defer db.Close()

	if match {
		c.Redirect(http.StatusMovedPermanently, "/dir?path=/")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func createaccount(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("password")
	
	db, err := sql.Open("sqlite3", Conf.DB)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	hasher := sha256.New()
	hasher.Write([]byte(pass))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	
	_, err = db.Exec("insert into Users values (null, ?, ?);", name, sha)
	
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func back(c *gin.Context) {
	fmt.Println("back called")
	b_body, _ := io.ReadAll(c.Request.Body)
	body := string(b_body)
	fmt.Println(body)

	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:8080/dir?path=/",
	})
}
