package bangumi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"bgm-calendar/util/http"
)

const (
	subjectTypeBook  subjectType = 1
	subjectTypeAnime subjectType = 2
	subjectTypeMusic subjectType = 3
	subjectTypeGame  subjectType = 4
	subjectTypeReal  subjectType = 5

	collectionTypeToWatch   collectionType = 1
	collectionTypeWatched   collectionType = 2
	collectionTypeWatching  collectionType = 3
	collectionTypeSuspended collectionType = 4
	collectionTypeAbandoned collectionType = 5
)

type Collections struct {
	Data  []Collection `json:"data"`
	Total int          `json:"total"`
	Limit int          `json:"limit"`
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

type subjectType int
type collectionType int

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
	limit := 50
	collections, err := getCollectionsByUsernameByPage(username, limit, 0)
	if err != nil {
		return Collections{}, err
	}
	if collections.Total <= limit {
		return collections, nil
	}
	batches := (collections.Total - 1) / limit
	wg := sync.WaitGroup{}
	wg.Add(batches)
	var errs []error
	for i := 0; i < batches; i++ {
		go func(offset int) {
			defer wg.Done()
			subCollections, err := getCollectionsByUsernameByPage(username, limit, offset)
			if err != nil {
				errs = append(errs, err)
				return
			}
			collections.Data = append(collections.Data, subCollections.Data...)
		}((i + 1) * limit)
	}
	wg.Wait()
	if len(errs) > 0 {
		return Collections{}, errs[0]
	}
	return collections, nil
}

func getCollectionsByUsernameByPage(username string, limit, offset int) (Collections, error) {
	u := url.URL{
		Scheme: "https",
		Host:   getAPIHost(),
		Path:   fmt.Sprintf("/v0/users/%s/collections", username),
	}
	query := url.Values{
		"subject_type": []string{fmt.Sprintf("%d", subjectTypeGame)},
		"type":         []string{fmt.Sprintf("%d", collectionTypeToWatch)},
		"limit":        []string{fmt.Sprintf("%d", limit)},
		"offset":       []string{fmt.Sprintf("%d", offset)},
	}
	u.RawQuery = query.Encode()
	client := http.NewHTTPClient()
	data, err := client.Get(u.String(), getHeaders())
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
	apiHost := os.Getenv("BANGUMI_API_HOST")
	if apiHost == "" {
		apiHost = "api.bgm.tv"
	}
	return apiHost
}

func getHeaders() map[string]string {
	headers := map[string]string{
		"User-Agent": "keo/bgm-calendar/0.0.1alpha",
	}
	accessToken := os.Getenv("BANGUMI_ACCESS_TOKEN")
	if accessToken != "" {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	}
	return headers
}
