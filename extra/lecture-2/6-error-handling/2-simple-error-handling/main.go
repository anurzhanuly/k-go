package main

import (
	"fmt"
	"log"
	"os"
)

// Чтение из файла с возвратом ошибки без её обработки

func main() {
	file, err := getFile("../README.md")
	if err != nil {
		log.Fatalln("не удалось прочитать README", err)
	}

	file, err = getFile("main.go")
	if err != nil {
		log.Fatalln("не удалось прочитать main.go", err)
	}

	fmt.Println(file)
}

func getFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		// ---------- возврат ошибки ----------
		return "", err
	}

	return string(file), nil
}
