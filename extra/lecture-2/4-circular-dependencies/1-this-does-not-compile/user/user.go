package user

import "go.kolesa-team.org/gl/go-course/extra/lecture-2/4-circular-dependencies/1-this-does-not-compile/advert"

type User struct {
	Id   int
	Name string
}

// Возвращает объявления пользователя
func GetUserAdverts(id int) []advert.Advert {
	return []advert.Advert{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
	}
}
