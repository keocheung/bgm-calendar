package bangumi

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"bgm-calendar/util/http"
)

const APIHost = "https://api.bgm.tv"

type Collections struct {
	Data []Collection `json:"data"`
}

type Collection struct {
	Subject Subject `json:"subject"`
}

type Subject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Date Date   `json:"date"`
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	s := string(b)
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "\"")
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = date
	return nil
}

func GetCollectionsByUsername(username string) (Collections, error) {
	client := http.NewHTTPClient()
	data, err := client.Get(fmt.Sprintf("https://%s/v0/users/%s/collections?subject_type=4&type=1&limit=50&offset=0", getAPIHost(), username), map[string]string{
		"User-Agent": "keo/bgm-calendar/0.0.1alpha",
	})
	if err != nil {
		return Collections{}, err
	}
	var collections Collections
	if err := json.Unmarshal(data, &collections); err != nil {
		return Collections{}, err
	}
	return collections, nil
}

func getAPIHost() string {
	apiHost := os.Getenv("BGM_CALENDAR_API_HOST")
	if apiHost == "" {
		apiHost = "api.bgm.tv"
	}
	return apiHost
}
