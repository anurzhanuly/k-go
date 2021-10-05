package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"go.kolesa-team.org/gl/toolkit/logger"
	"net/http"
	"runtime"
)

// Config структура конфигов
type Config struct {
	Listen string         // Адрес и порт
	Debug  bool           // Режим отладки
	Logger logger.Options // Опции логирования
}

// main точка входа
func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	cfg := new(Config)
	if _, err := toml.DecodeFile(*configPath, cfg); err != nil {
		logrus.
			WithField("config", *configPath).
			WithError(err).
			Fatalln("Невозможно загрузить конфигурацию сервиса")
	}

	logger.Init(cfg.Logger, cfg.Debug)

	mux := http.NewServeMux()

	mux.HandleFunc("/release", showRelease)
	mux.HandleFunc("/health", health)

	logrus.Info("Сервис запущен")

	if err := http.ListenAndServe(cfg.Listen, mux); err != nil {
		logrus.WithError(err).Error("Сервис остановлен с ошибкой")
	}
}

// health используется для проверки состояния микросервиса
func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

// showRelease возвращает информацию о релизе
func showRelease(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf8")
	_, _ = w.Write([]byte(`{
		"revision": "rev-unset",
		"branch": "branch-unset",
		"build": "build-unset",
		"buildUrl": "result-unset",
		"go":  "` + runtime.Version() + `",
		"date": "date-unset"
	}`))
}
