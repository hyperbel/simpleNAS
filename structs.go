package main

type Config struct {
	Dir	string	`json:"dir"`
}

type FileInfo struct {
	Name string
	IsDir bool
	Size int
}
