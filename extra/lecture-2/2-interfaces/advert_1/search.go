package advert_1

import "go.kolesa-team.org/gl/go-course/extra/lecture-2/2-interfaces/api_client"

// Возвращает объявления по запросу
func FindAdvertsByQuery(query string) []string {
	// сходим в эластик и соберём id подходящих объявлений
	ids := []int{1, 2, 3}

	// сходим в api за объявлениями
	apiClient := api_client.Client{}
	adverts := []string{}

	for _, id := range ids {
		advert := apiClient.GetAdvert(id)
		adverts = append(adverts, advert)
	}

	return adverts
}
