package main

import (
	"fmt"
	"net/http"
)

var lastName = ""

func main() {
	fmt.Println("Программа на Go запущена")

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/history", historyHandler)
	http.ListenAndServe(":8080", nil)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", lastName)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	lastName = r.URL.Query().Get("name")

	// пишем клиенту
	fmt.Fprintf(w, "Салем, %s!\n", lastName)

	// пишем в stdout
	fmt.Println("запрос обработан")
}
