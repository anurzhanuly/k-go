package infrastructure

import (
	"github.com/BurntSushi/toml"
	tracerCfg "github.com/uber/jaeger-client-go/config"
	"go.kolesa-team.org/gl/microservice/db"
	"go.kolesa-team.org/gl/toolkit/logger"
	"go.kolesa-team.org/gl/toolkit/stats"
)

// Config конфигурация сервиса
type Config struct {
	Listen string                  // Адрес и порт
	Debug  bool                    // Режим отладки
	Logger logger.Options          // Опции логирования
	Stats  stats.Options           // Опции сбора метрик
	Tracer tracerCfg.Configuration // Опции трассировщика запросов
	Mongo  db.Options              // Опции для монго
}

// NewConfig возвращает новую конфигурацию
func NewConfig() Config {
	return Config{}
}

// FromFile инициализирует конфигурацию из файла
func (c *Config) FromFile(file string) error {
	_, err := toml.DecodeFile(file, c)
	return err
}
