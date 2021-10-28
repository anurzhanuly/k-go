package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type FileReadErr struct {
	Path        string
	OriginalErr error
}

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
		logrus.WithError(err).WithField("path", path).Error("не удалось прочитать файл")
		return "", &FileReadErr{OriginalErr: err, Path: path}
	}

	return string(file), nil
}
