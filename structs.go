package main

type Config struct {
	Dir string `json:"dir"`
	DB  string `json:"db"`
}

type FileInfo struct {
	Name  string
	IsDir bool
	Size  int
}

type User struct {
	id     int
	name   string
	passwd string
}

type RemoveFilesRequestBody struct {
	Files  []string `json:"files"`
	Search string   `json:"search"`
}

type CreateDirRequestBody struct {
	Search string `json:"search"`
}
