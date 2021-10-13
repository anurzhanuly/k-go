package advert

import "go.kolesa-team.org/ba/go-course/extra/lecture-2/4-circular-dependencies/3-many-packages/user"

type Advert struct {
	Id   int
	User user.User
}
