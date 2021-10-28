package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

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
		return "", errors.Wrapf(err, "не удалось прочитать файл %s", path)
	}

	return string(file), nil
}
