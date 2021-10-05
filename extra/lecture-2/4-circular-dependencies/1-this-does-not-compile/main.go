package main

import (
	"fmt"
	"go.kolesa-team.org/gl/go-course/extra/lecture-2/4-circular-dependencies/1-this-does-not-compile/user"
)

// Программа не скомпилируется из-за циклических зависимостей:
// user зависит от advert, а advert - от user
func main() {
	fmt.Println(user.GetUserAdverts(1))
}
