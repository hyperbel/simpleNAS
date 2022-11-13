package main

import (
	"os"
//	"fmt"
	"log"
	"net/http"
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
	c.HTML(http.StatusOK, "dir.html", gin.H{
		"dir": Conf.Dir,
		"files": fs,
	})

}

/*
func path(c *gin.Context) {
	fmt.Println(Conf.Dir)
	fmt.Println(c.Param("path"))
	dir := Conf.Dir + c.Param("path")
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		c.HTML(500, "error.html", gin.H{
			"message": "An error occurred, please check logs",
		})
	}
	
	fs := make([]FileInfo, len(files))
	
	for index, file := range files {
		fs[index] = FileInfo{file.Name(), file.IsDir(), 0}
	}
	
	c.HTML(http.StatusOK, "dir.html", gin.H{
		"dir": dir,
		"files": fs,
	})
}
*/

func dir(c *gin.Context) {
	path := c.Query("path")

	if path == "" {
		c.HTML(404, "error.html", gin.H{
			"message": "redirecting...",
		})
	}
	
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
