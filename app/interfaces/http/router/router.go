package router

import (
	"go.kolesa-team.org/gl/microservice/app/infrastructure"
	"net/http"
)

// NewRouter возвращает новый роутер приложения
func NewRouter(container infrastructure.Container, debug bool) *http.ServeMux {
	mux := http.NewServeMux()

	// Сюда записываем роуты

	return mux
}
