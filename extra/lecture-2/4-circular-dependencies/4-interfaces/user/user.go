package user

import "go.kolesa-team.org/ba/go-course/extra/lecture-2/4-circular-dependencies/4-interfaces/advert"

type User struct {
	id   int
	name string
}

func (u User) Id() int {
	return u.id
}

func (u User) Name() string {
	return u.name
}

// Возвращает объявления пользователя
func GetUserAdverts(id int) []advert.Advert {
	return []advert.Advert{
		{
			Id:   1,
			User: User{name: "Rob"},
		},
		{
			Id:   2,
			User: User{name: "Bob"},
		},
	}
}
