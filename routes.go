package main

import (
	"fmt"
	"os"
	"log"
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
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	username := c.PostForm("uname")
	password := c.PostForm("passwd")
	fmt.Println(username, password)

	hash := sha256.New()
	pwh := hash.Sum([]byte(password))[:]
	fmt.Println(string(pwh[:]))

	q := fmt.Sprintf("select * from Users where name=\"%s\" and password=0;", username)//, pwh)
	fmt.Println(q)
	var u User
	err = db.QueryRow(q, username, pwh).Scan(&u.id, &u.name, &u.passwd)
	match := false
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {

		match = true
	}
	
	/*
	for rows.Next() {
		var p []byte
		err = rows.Scan(&p)

		if err != nil {
			log.Fatal(err)
		} else {
			match = true
			break
		}
	}
*/
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
	fmt.Println(sha)

	
	ret, err := db.Exec("insert into Users values (null, ?, ?);", name, sha)
	fmt.Println(ret)
	
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
