package main

import (
	"log"
	"net/http"
	"os"

	"bgm-calendar/controller"
	"bgm-calendar/meta"
	"bgm-calendar/util/logger"
)

func main() {
	logger.Infof("bgm-calendar %s", meta.Version)
	http.HandleFunc("/users/", controller.Games)
	err := http.ListenAndServe(":"+getPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := os.Getenv("BGM_CALENDAR_PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
