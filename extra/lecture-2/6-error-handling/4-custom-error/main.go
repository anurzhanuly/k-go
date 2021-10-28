package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// Использование кастомных структур с информацией об ошибках

// Ошибка чтения файла
type FileReadErr struct {
	// По какому пути пытались прочитать файл
	Path string
	// Исходная ошибка, возвращённая функцией os.ReadFile
	OriginalErr error
}

// Возвращает описание ошибки.
// Благодаря этой функции FileReadErr реализует интерфейс error.
func (f FileReadErr) Error() string {
	return fmt.Sprintf("не удалось прочитать файл %s: %s", f.Path, f.OriginalErr)
}

func main() {
	file, err := getFile("../README.md")
	if err != nil {
		log.Fatalln(err)
	}

	file, err = getFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(file)
}

func getFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		// логируем ошибку
		logrus.WithError(err).WithField("path", path).Error("не удалось прочитать файл")

		// возвращаем ошибку кастомного типа
		return "", &FileReadErr{OriginalErr: err, Path: path}
	}

	return string(file), nil
}
