package advert

type User interface {
	GetId() int
	GetName() string
}

type Advert struct {
	Id   int
	User User
}
