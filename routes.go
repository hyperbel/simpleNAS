package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "index called",
	})
}

func dirtest(c *gin.Context) {
	files, err := ioutil.ReadDir(Conf.Dir)
	fmt.Println(Conf.Dir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var fs [100]FileInfo	//change 100 to amount of files
	
	for i, f := range files {
		fs[i] = FileInfo{f.Name(), f.IsDir(), 0}
		fmt.Println(Conf.Dir, f.Name(), f.IsDir())
	}
	c.HTML(http.StatusOK, "dirtest.html", gin.H{
		"message": "directory test stuff",
		"files": fs,
	})

}
