package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"bgm-calendar/controller"
	"bgm-calendar/meta"
	"bgm-calendar/util/logger"
)

func main() {
	logger.Infof("bgm-calendar %s (BuildTime: %s)", meta.Version, meta.BuildTime)
	http.HandleFunc("/users/", controller.Users)
	port := getPort()
	logger.Infof("bgm-calendar listening on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
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
