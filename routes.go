package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// login / signup page
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

// get current directory
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

	url := location.Get(c)
	fmt.Printf("%+v", url.RawPath)
	fmt.Println("test")

	c.HTML(http.StatusOK, "dir.html", gin.H{
		"dir":      dir,
		"files":    fs,
		"userid":   uid,
		"location": "not working",
	})
}

// logs the user in, if passwd and username are correct
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

// allows the user to create an account
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

// not fully implemented "back" function for going up in dir structure
func back(c *gin.Context) {
	b_body, _ := io.ReadAll(c.Request.Body)
	body := string(b_body)

	c.JSON(http.StatusOK, gin.H{
		"url": body,
	})
}

// creates directory on users command
func createdir(c *gin.Context) {
	dir := c.Query("name")
	fmt.Println(dir)
	var create_dir_request_body CreateDirRequestBody
	err := c.BindJSON(&create_dir_request_body)
	handleError(err, 0)
	path := pathFromQuery(create_dir_request_body.Search)
	fmt.Println(path)
	full_path := fmt.Sprintf("%s/%s", path, dir)
	err = os.Mkdir(full_path, os.ModePerm)
	handleError(err, 1)
	c.JSON(http.StatusOK, gin.H{})
}

// executes, when the user tries to remove a file
func removefiles(c *gin.Context) {
	var remove_files_request_body RemoveFilesRequestBody
	err := c.BindJSON(&remove_files_request_body)
	handleError(err, 0)

	path := pathFromQuery(remove_files_request_body.Search)

	for _, file_name := range remove_files_request_body.Files {
		for i := len(file_name) - 1; i > -1; i-- {
			if file_name[i] == 95 {
				file := file_name[:len(file_name)-(len(file_name)-i)]
				full_path := fmt.Sprintf("%v%v", path, file)
				fmt.Println(full_path)
				if _, err := os.Stat(full_path); errors.Is(err, os.ErrNotExist) {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": fmt.Sprintf("file %v does not exist!", full_path),
					})
				} else {
					eerr := os.RemoveAll(full_path)
					handleError(eerr, 1)
				}
			}
		}
	}
}

// gets executed when user attempts to upload a file
func uploadfile(c *gin.Context) {
	fmt.Println("getting from submitted stuff")
	fmt.Println(c.PostForm("hidden_url"))
	path := pathFromQuery(c.PostForm("hidden_url"))
	file, err := c.FormFile("file_upload")
	handleError(err, 1)

	fmt.Println("settings vars")
	ext := filepath.Ext(file.Filename)
	fullFileName := path + file.Filename + ext
	fmt.Println(fullFileName)

	fmt.Println("creating file")
	os_file, err := os.Create(fullFileName)
	handleError(err, 1)
	fmt.Println("os file created successfully")
	os_file.Close()

	err = c.SaveUploadedFile(file, fullFileName)
	handleError(err, 1)

	c.Redirect(301, "/dir?path=/") // 301 is required to redirect as GET instead of POST when using 308
}
