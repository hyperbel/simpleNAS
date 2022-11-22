package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
)

func index(c *gin.Context) {
	files, err := os.ReadDir(Conf.Dir)
	handleError(err, 1)

	fs := make([]FileInfo, len(files)) //change 100 to amount of files

	for i, f := range files {
		fs[i] = FileInfo{f.Name(), f.IsDir(), 0}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"dir":   Conf.Dir,
		"files": fs,
	})

}

func dir(c *gin.Context) {
	path := c.Query("path")
	//db, err := sql.Open("sqlite3", Conf.DB)
	session := sessions.Default(c)

	uid := session.Get("userid")
	fmt.Println(uid)

	dir := Conf.Dir + path
	files, err := os.ReadDir(dir)

	handleError(err, 1)

	fs := make([]FileInfo, len(files))

	for i, file := range files {
		fs[i] = FileInfo{file.Name(), file.IsDir(), 0}
	}
	hist := session.Get("history")
	if hist == "" {
		session.Set("history", dir+",")
		fmt.Println(session.Get("history"))
	} else {
		session.Set("history", fmt.Sprintf("%v,%v", hist, dir))
		fmt.Println(session.Get("history"))
	}
	session.Save()

	c.HTML(http.StatusOK, "dir.html", gin.H{
		"dir":    dir,
		"files":  fs,
		"userid": uid,
	})
}

func login(c *gin.Context) {
	db, err := sql.Open("sqlite3", Conf.DB)
	session := sessions.Default(c)
	hasher := sha256.New()
	var match bool
	var u User

	handleError(err, 1)

	username := c.PostForm("uname")
	password := c.PostForm("passwd")

	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	q, err := db.Prepare("SELECT * FROM Users WHERE name=? and password=?")
	handleError(err, 0)

	err = q.QueryRow(username, sha).Scan(&u.id, &u.name, &u.passwd)

	if err != nil {
		match = false
		log.Println(err)
	}
	match = true

	defer db.Close()

	if match {
		session.Set("userid", u.id)
		session.Set("history", "")
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/dir?path=")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func createaccount(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("password")

	db, err := sql.Open("sqlite3", Conf.DB)
	handleError(err, 1)

	hasher := sha256.New()
	hasher.Write([]byte(pass))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	_, err = db.Exec("insert into Users values (null, ?, ?);", name, sha)

	handleError(err, 1)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func back(c *gin.Context) {
	b_body, _ := io.ReadAll(c.Request.Body)
	body := string(b_body)

	c.JSON(http.StatusOK, gin.H{
		"url": body,
	})
}

func createdir(c *gin.Context) {
	dir := c.Query("name")
	err := os.Mkdir(fmt.Sprintf("%s/%s", Conf.Dir, dir), os.ModePerm)
	handleError(err, 1)
	c.JSON(http.StatusOK, gin.H{})
}

func removefiles(c *gin.Context) {
	var remove_files_request_body RemoveFilesRequestBody
	err := c.BindJSON(&remove_files_request_body)
	handleError(err, 0)
	for _, file_name := range remove_files_request_body.Files {
		for i := len(file_name) - 1; i > -1; i-- {
			if file_name[i] == 95 {
				fmt.Printf("%v : %v : %v\n", i, string(file_name[i]), file_name[i])
				file := file_name[:len(file_name)-(len(file_name)-i)]
				fmt.Printf("%+v\n", file)
			}
		}
	}
}
