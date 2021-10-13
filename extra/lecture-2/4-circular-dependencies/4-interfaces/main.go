package main

import (
	"fmt"
	"go.kolesa-team.org/ba/go-course/extra/lecture-2/4-circular-dependencies/4-interfaces/user"
)

func main() {
	fmt.Println(user.GetUserAdverts(1))
}
