package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"go.kolesa-team.org/gl/microservice/app/infrastructure"
	"go.kolesa-team.org/gl/toolkit/jaeger"
	"go.kolesa-team.org/gl/toolkit/logger"
	"go.kolesa-team.org/gl/toolkit/stats"
	"net/http"
	"runtime"
)

// main точка входа
func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	cfg := infrastructure.NewConfig()
	if err := cfg.FromFile(*configPath); err != nil {
		logrus.
			WithField("config", *configPath).
			WithError(err).
			Fatalln("Невозможно загрузить конфигурацию сервиса")
	}

	logger.Init(cfg.Logger, cfg.Debug)

	if _, err := jaeger.Init(cfg.Tracer); err != nil {
		logrus.WithError(err).Error("Не удалось инициализировать трассировку")
	}

	if err := stats.Init(cfg.Stats); err != nil {
		logrus.WithError(err).Error("Не удалось инициализировать статистику")
	}

	_ = infrastructure.NewContainer(cfg)

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
