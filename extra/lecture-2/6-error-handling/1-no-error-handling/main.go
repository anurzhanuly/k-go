package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(getFile())
}

func getFile() string {
	file, _ := os.ReadFile("../README.md")

	return string(file)
}
