package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Lines(filename string) []string {
	fileBytes, err := os.ReadFile("./inputs/" + filename + ".txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return strings.Split(string(fileBytes), "\n")
}
