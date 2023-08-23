package main

import (
	"log"
	"net/http"

	"bgm-calendar/controller"
)

func main() {
	http.HandleFunc("/users/", controller.Games)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
