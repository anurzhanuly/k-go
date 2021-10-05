package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Config структура конфигов
type Config struct {
	Listen   string // Адрес и порт
	LogLevel string // Уровень логирования
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

	lvl, _ := logrus.ParseLevel(cfg.LogLevel)

	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	mux := http.NewServeMux()

	mux.HandleFunc("/health", health)

	logrus.Info("Сервис запущен")

	if err := http.ListenAndServe(cfg.Listen, mux); err != nil {
		logrus.WithError(err).Error("Сервис остановлен с ошибкой")
	}
}

// health используется для проверки состояния микросервиса
func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("ok"))
}
