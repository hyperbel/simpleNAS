package main

type Config struct {
	Dir	string	`json:"dir"`
}

type FileInfo struct {
	Name string
	IsDir bool
	Size int
}

type User struct {
	Id int
	Name string `form:"uname" binding:"required"`
	PasswordHash byte `form:"passwd" binding:"required"`
}
