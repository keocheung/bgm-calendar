package controller

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"bgm-calendar/pkg/bangumi"
)

const (
	calendarPrefix = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:bgm-calendar
CALSCALE:GREGORIAN
X-WR-CALNAME:想玩的游戏
X-APPLE-LANGUAGE:zh
X-APPLE-REGION:CN
`
	calendarSuffix = "END:VCALENDAR"
	eventTemplete  = `BEGIN:VEVENT
DTSTAMP;VALUE=DATE:19760401
UID:%s
DTSTART;VALUE=DATE:%s
CLASS:PUBLIC
SUMMARY;LANGUAGE=zh_CN:%s
TRANSP:TRANSPARENT
CATEGORIES:想玩的游戏
END:VEVENT
`
)

var gamesPathPattern = regexp.MustCompile("^/users/(.+)/games.ics$")

func Games(w http.ResponseWriter, r *http.Request) {
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
	var events []string
	preferCNConfig := os.Getenv("BGM_CALENDAR_PREFER_CN_NAME")
	preferCN := preferCNConfig == "true" || preferCNConfig == "1"
	startTime := time.Now().AddDate(0, -30, 0)
	for _, collection := range collections.Data {
		if collection.Subject.Date.Before(startTime) {
			continue
		}
		date := collection.Subject.Date
		event := fmt.Sprintf(eventTemplete, fmt.Sprintf("BANGUMI-SUBJECT-%d", collection.Subject.Id),
			date.Format("20060102"), getSubjectName(collection.Subject, preferCN))
		events = append(events, event)
	}
	cal := calendarPrefix
	for _, event := range events {
		cal += event
	}
	cal += calendarSuffix
	w.Write([]byte(cal))
}

func getSubjectName(subject bangumi.Subject, preferCN bool) string {
	if subject.NameCN != "" && preferCN {
		return subject.NameCN
	}
	return subject.Name
}
