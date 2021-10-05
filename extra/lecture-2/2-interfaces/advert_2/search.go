package advert_2

// Возвращает объявления по запросу
func FindAdvertsByQuery(query string, apiClient ApiClientInterface) []string {
	// сходим в эластик и соберём id подходящих объявлений
	ids := []int{1, 2, 3}

	// сходим в api за объявлениями
	adverts := []string{}

	for _, id := range ids {
		advert := apiClient.GetAdvert(id)
		adverts = append(adverts, advert)
	}

	return adverts
}
