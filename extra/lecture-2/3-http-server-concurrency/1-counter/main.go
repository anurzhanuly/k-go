package main

import (
	"fmt"
	"net/http"
)

// Количество обработанных программой запросов
var requestCount = 0

func main() {
	http.HandleFunc("/increment", incrementHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":8080", nil)
}

// Увеличиваем счётчик запросов
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Fprintf(w, "%d\n", requestCount)
}

// Вывводим количество запросов
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d\n", requestCount)
}
