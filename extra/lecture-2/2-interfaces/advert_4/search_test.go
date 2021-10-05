package advert_1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Фейковая реализация хранилища объявлений для тестов.
// Требуется реализовать только 1 метод, а не все методы, объявленные в api_client.
type fakeAdvertStorage struct{}

func (f fakeAdvertStorage) GetAdvert(id int) string {
	if id == 1 {
		return "advert-a"
	} else {
		return "advert-b"
	}
}

// Проверяем поиск объявлений, подставив тестовую реализацию хранилища объявлений.
func TestFindAdvertsByQuery(t *testing.T) {
	finder := NewAdvertFinder(fakeAdvertStorage{})

	actualAdverts := finder.FindAdvertsByQuery("my-query")
	expectedAdverts := []string{"advert-a", "advert-b"}
	require.Equal(t, actualAdverts, expectedAdverts)
}
