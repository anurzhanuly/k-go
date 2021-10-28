package user_adverts

import (
	"go.kolesa-team.org/ba/go-course/extra/lecture-2/4-circular-dependencies/3-many-packages/advert"
	"go.kolesa-team.org/ba/go-course/extra/lecture-2/4-circular-dependencies/3-many-packages/user"
)

// Возвращает объявления пользователя
func GetUserAdverts(id int) []advert.Advert {
	return []advert.Advert{
		{
			Id:   100,
			User: user.User{Name: "Azatbek"},
		},
		{
			Id:   200,
			User: user.User{Name: "Vitaliy"},
		},
	}
}
