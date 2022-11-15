package main

type Config struct {
	Dir	string	`json:"dir"`
	DB string `json:"db"`
}

type FileInfo struct {
	Name string
	IsDir bool
	Size int
}
