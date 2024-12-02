package main

import (
	"log"
	"os"
	"path/filepath"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, perm)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func main() {
	err := WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}
}