package main

import (
	"fmt"
	"go.kolesa-team.org/ba/go-course/extra/lecture-2/4-circular-dependencies/2-single-package/stuff"
)

func main() {
	fmt.Println(stuff.GetUserAdverts(1))
}
