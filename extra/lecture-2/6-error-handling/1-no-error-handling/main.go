package main

import (
	"fmt"
	"os"
)

// Чтение из файла с игнорированием ошибки

func main() {
	fmt.Println(getFile())
}

func getFile() string {
	file, _ := os.ReadFile("../README.md")

	return string(file)
}
