package advert_1

// Объявляем интерфейс нашей зависимости
// Здесь неважно, кто именно её реализует: клиент API или иной пакет.
type AdvertStorage interface {
	GetAdvert(id int) string
}

// Возвращает объявления по запросу
func FindAdvertsByQuery(query string, advertStorage AdvertStorage) []string {
	// сходим в эластик и соберём id подходящих объявлений
	ids := []int{1, 2}

	// сходим в api за объявлениями
	adverts := []string{}

	for _, id := range ids {
		advert := advertStorage.GetAdvert(id)
		adverts = append(adverts, advert)
	}

	return adverts
}
