package advert_1

// Объявляем интерфейс нашей зависимости
// Здесь неважно, кто именно её реализует: клиент API или иной пакет.
type AdvertStorage interface {
	GetAdvert(id int) string
}

// Сервис поиска объявлений.
// Структура хранит внутри свои зависимости
// (помимо хранилища объявлений здесь же может быть elasticsearch, кэш и т.д.)
type AdvertFinder struct {
	storage AdvertStorage
}

// Создаёт сервис поиска объявлений.
func NewAdvertFinder(storage AdvertStorage) AdvertFinder {
	return AdvertFinder{storage: storage}
}

// Возвращает объявления по запросу
func (finder AdvertFinder) FindAdvertsByQuery(query string) []string {
	// сходим в эластик и соберём id подходящих объявлений
	ids := []int{1, 2}

	// сходим в api за объявлениями
	adverts := []string{}

	for _, id := range ids {
		advert := finder.storage.GetAdvert(id)
		adverts = append(adverts, advert)
	}

	return adverts
}
