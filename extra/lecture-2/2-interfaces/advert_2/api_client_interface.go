package advert_2

// Создаём интерфейс клиента API также, как обычно делаем в PHP:
// описываем все методы, даже если в текущем контексте нам нужен какой-то один.
// Чтобы подменить такой интерфейс в тестах, придётся писать реализацию каждого метода (
type ApiClientInterface interface {
	GetUser(id int) string
	CreateUser(id int)
	GetAdvert(id int) string
	GetCategory(id int) string
	GetParameter(id int) string
}
