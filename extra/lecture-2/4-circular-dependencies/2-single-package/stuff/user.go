package stuff

type User struct {
	Id   int
	Name string
}

// Возвращает объявления пользователя
func GetUserAdverts(id int) []Advert {
	return []Advert{
		{
			Id:   1,
			User: User{Name: "Rob"},
		},
		{
			Id:   2,
			User: User{Name: "Bob"},
		},
	}
}
