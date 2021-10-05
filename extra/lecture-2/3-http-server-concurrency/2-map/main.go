package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Хэш-таблица (ассоциативный массив) вида "дата запроса => IP адрес клиента"
var userIpsByDate = map[time.Time]string{}

func main() {
	http.HandleFunc("/request", requestHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":8080", nil)
}

// Запоминаем дату и IP адрес клиента при обращении по адресу /request
func requestHandler(w http.ResponseWriter, r *http.Request) {
	userIpsByDate[time.Now()] = r.RemoteAddr

	fmt.Fprintf(w, "hello!\n")
}

// Вывводим статистику запросов
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	datesJson, _ := json.MarshalIndent(userIpsByDate, "", "  ")
	fmt.Fprintf(w, "%s\n", datesJson)
}
