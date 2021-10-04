package infrastructure

import (
	"github.com/sirupsen/logrus"
	"go.kolesa-team.org/gl/microservice/db"
)

// Container контейнер для хранения зависимостей
type Container struct {
	Config Config
}

// NewContainer возвращает новый контейнер зависимостей
func NewContainer(cfg Config) Container {
	_, err := db.Dial(cfg.Mongo)
	if err != nil {
		logrus.WithError(err).
			WithField("url", cfg.Mongo.URL).Fatalln("Не удалось инициализировать соединение с mongodb")
	}

	return Container{
		Config: cfg,
	}
}
