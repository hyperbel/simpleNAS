package main

import (
	"fmt"
	"os"
)

func parseHistoryFromString(hist string) []string {

	return make([]string, 0)
}

func parseHistoryToString(hist []string) string {
	return ""
}

func handleArgs(args []string) string {
	config_file_location := ""
	home_dir, err := os.UserHomeDir()

	if len(args) == 1 {
		handleError(err, 1)
		config_file_location = fmt.Sprintf("%s/.config/simplenas/config.json", home_dir)
		fmt.Println(config_file_location)
		return config_file_location
	}

	if args[1] == "help" {
		fmt.Println("You can provide a config file, by passing it as an Argument.")
		fmt.Println("Usage: go run . <config>")
		fmt.Println("the default config file is ~/.config/simplenas/config")
		fmt.Println("If this config file is not there, you have to provide one")
		os.Exit(0)
	} else {
		config_file_location = args[1]
		fmt.Println(args[1])
	}

	return config_file_location
}

func handleError(e error, exit_code int) {
	if e != nil {
		fmt.Println(e)
		if exit_code != 0 {
			os.Exit(exit_code)
		}
	}
}

func pathFromQuery(query string) string {
	path := query[6:]
	return Conf.Dir + path
}
