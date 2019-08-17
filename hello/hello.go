package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func files() ([]string, error) {
	searchDir := "C:/Users/<path>"
	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		fmt.Println(file)
	}
	return fileList, nil
}

func main() {
	files()
}
