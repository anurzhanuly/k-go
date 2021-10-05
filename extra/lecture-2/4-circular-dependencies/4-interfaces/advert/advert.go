package advert

type User interface {
	Id() int
	Name() string
}

type Advert struct {
	Id   int
	User User
}
