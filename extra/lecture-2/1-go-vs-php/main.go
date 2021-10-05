package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Программа на Go запущена")

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// пишем клиенту
	fmt.Fprintf(w, "Привет, %s!\n", name)

	// пишем в stdout
	fmt.Println("запрос обработан")
}
