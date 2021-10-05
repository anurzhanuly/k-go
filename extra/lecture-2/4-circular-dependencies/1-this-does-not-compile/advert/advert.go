package advert

import "go.kolesa-team.org/gl/go-course/extra/lecture-2/4-circular-dependencies/1-this-does-not-compile/user"

type Advert struct {
	Id   int
	User user.User
}
