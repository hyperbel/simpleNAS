package main

import (
	"os"
	"log"
	"net/http"
	"database/sql"
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
	
	c.JSON(http.StatusOK, gin.H{
		"lastDir": ,
	})
}
