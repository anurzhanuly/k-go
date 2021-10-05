package user_adverts

import (
	"go.kolesa-team.org/gl/go-course/extra/lecture-2/4-circular-dependencies/3-many-packages/advert"
	"go.kolesa-team.org/gl/go-course/extra/lecture-2/4-circular-dependencies/3-many-packages/user"
)

// Возвращает объявления пользователя
func GetUserAdverts(id int) []advert.Advert {
	return []advert.Advert{
		{
			Id:   1,
			User: user.User{Name: "Rob"},
		},
		{
			Id:   2,
			User: user.User{Name: "Bob"},
		},
	}
}
