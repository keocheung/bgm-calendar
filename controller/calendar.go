package controller

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"bgm-calendar/meta"
	"bgm-calendar/pkg/bangumi"

	ics "github.com/arran4/golang-ical"
)

const (
	calendarGames = "想玩的游戏"
)

var gamesPathPattern = regexp.MustCompile("^/users/(.+)/games.ics$")

func Users(w http.ResponseWriter, r *http.Request) {
	matches := gamesPathPattern.FindStringSubmatch(r.URL.Path)
	if matches == nil {
		http.NotFound(w, r)
		return
	}
	username := matches[1]
	collections, err := bangumi.GetCollectionsByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cal := generateGamesCal(collections.Data)
	w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
	w.Write([]byte(cal))
}

func generateGamesCal(collections []bangumi.Collection) string {

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetName(calendarGames)
	cal.SetProductId(meta.UserAgent)
	preferCNConfig := os.Getenv("BGM_CALENDAR_PREFER_CN_NAME")
	preferCN := preferCNConfig == "1" || preferCNConfig == "true"
	startTime := time.Now().AddDate(0, -1, 0)
	for _, collection := range collections {
		if collection.Subject.Date.Before(startTime) {
			continue
		}
		event := cal.AddEvent(fmt.Sprintf("BANGUMI-SUBJECT-%d", collection.Subject.Id))
		event.SetAllDayStartAt(collection.Subject.Date.Time)
		event.SetSummary(getSubjectName(collection.Subject, preferCN))
		event.SetProperty(ics.ComponentPropertyCategories, calendarGames)
	}
	return cal.Serialize()
}

func getSubjectName(subject bangumi.Subject, preferCN bool) string {
	if subject.NameCN != "" && preferCN {
		return subject.NameCN
	}
	return subject.Name
}
