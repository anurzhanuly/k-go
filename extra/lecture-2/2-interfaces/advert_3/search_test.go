package advert_1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Фейковая реализация хранилища объявлений для тестов.
// Требуется реализовать только 1 метод, а не все методы, объявленные в api_client.
type fakeAdvertStorage struct{}

func (f fakeAdvertStorage) GetAdvert(id int) string {
	// Наделим тестовое хранилище какой-то примитивной логикой
	if id == 1 {
		return "advert-a"
	} else {
		return "advert-b"
	}
}

// Проверяем поиск объявлений, подставив тестовую реализацию хранилища объявлений.
func TestFindAdvertsByQuery(t *testing.T) {
	actualAdverts := FindAdvertsByQuery("my-query", fakeAdvertStorage{})
	expectedAdverts := []string{"advert-a", "advert-b"}
	require.Equal(t, actualAdverts, expectedAdverts)
}
