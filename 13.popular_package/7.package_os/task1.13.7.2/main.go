package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(data io.Reader, fd io.Writer) error {
	if file, ok := fd.(*os.File); ok {
		dir := filepath.Dir(file.Name())
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}
	}

	// Копируем данные из data в fd
	_, err := io.Copy(fd, data)
	return err
}

func main() {
	filePath := "course1/13.popular_package/7.package_os/task1.13.7.2/file.txt"
	// Открываем файл для записи
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close() // отложенная функция закрытия дескриптора файла

	err = WriteFile(strings.NewReader("Hello, World!"), file)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
	}
}