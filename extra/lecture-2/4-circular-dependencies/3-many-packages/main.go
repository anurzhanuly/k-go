package main

import (
	"fmt"
	"go.kolesa-team.org/gl/go-course/extra/lecture-2/4-circular-dependencies/3-many-packages/user_adverts"
)

func main() {
	fmt.Println(user_adverts.GetUserAdverts(1))
}
