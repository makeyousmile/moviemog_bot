package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	baseUrl         = "https://suggest-kinopoisk.yandex.net/suggest-kinopoisk?srv=kinopoisk&part=%s&_=%d"
	FilmURL         = "https://plus.kinopoisk.ru/film/%d/"
	PersonURL       = "https://plus.kinopoisk.ru/name/%d/"
	FilmIDRegexp    = regexp.MustCompile(`.*\ id=([0-9]+)COBJECT$`)
	PersonIDRegexpt = regexp.MustCompile(`.*\ id=([0-9]+)PERSON$`)
)

type Data struct {
	Query      string
	resultsRaw []string
	objectsRaw []string
	FilmIDs    []int
	PersonIDs  []int
	Films      []Film
	Persons    []Person
}

type Film struct {
	URL              string       `json:"-"`
	EntityId         int          `json:"entityId"`
	SearchObjectType string       `json:"searchObjectType"`
	OriginalTitle    string       `json:"originalTitle"`
	Title            string       `json:"title"`
	Years            []int        `json:"years"`
	IsCompleted      bool         `json:"isCompleted"`
	Type             string       `json:"type"`
	Genres           []string     `json:"genres"`
	Snippet          Image        `json:"snippet"`
	Poster           Image        `json:"poster"`
	Expectations     Expectations `json:"expectations"`
	Rating           Rating       `json:"rating"`
}

type Person struct {
	URL              string `json:"-"`
	EntityId         int    `json:"entityId"`
	SearchObjectType string `json:"searchObjectType"`
	OriginalName     string `json:"originalName"`
	Name             string `json:"name"`
	Age              int    `json:"age"`
	BirthDate        string `json:"birthDate"`
	BirthPlaceString string `json:"birthPlaceString"`
	ZodiacSign       string `json:"zodiacSign"`
	Roles            []Role `json:"roles"`
	Gender           string `json:"gender"`
	FirstFilmYear    int    `json:"firstFilmYear"`
	LatestFilmYear   int    `json:"latestFilmYear"`
	Picture          Image  `json:"picture"`
}

type Role struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Image struct {
	BaseUrl string `json:"baseUrl"`
}

type Expectations struct {
	Votes   int     `json:"votes"`
	Rate    float32 `json:"rate"`
	IsReady bool    `json:"idReady"`
	Ready   bool    `json:"ready"`
}

type Rating struct {
	Votes   int     `json:"votes"`
	Max     int     `json:"max"`
	Rate    float32 `json:"rate"`
	IsReady bool    `json:"idReady"`
	Ready   bool    `json:"ready"`
}

func (d *Data) UnmarshalJSON(data []byte) error {
	dataSlice := make([]json.RawMessage, 3)
	err := json.Unmarshal(data, &dataSlice)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataSlice[0], &d.Query)
	if err != nil {
		return err
	}
	if len(dataSlice) < 3 {
		return errors.New("Incorrent data format, missing result data")
	}
	err = json.Unmarshal(dataSlice[1], &d.resultsRaw)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataSlice[2], &d.objectsRaw)
	if err != nil {
		return err
	}
	return nil
}

func encodeUrl(query string) string {
	queryEnc := url.QueryEscape(query)
	return fmt.Sprintf(baseUrl, queryEnc, time.Now().UTC().Unix())
}

func do(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (d *Data) decode(body []byte) error {
	if err := json.Unmarshal(body, d); err != nil {
		return err
	}
	return nil
}

func (d *Data) processIDs() error {
	var res []string
	var id string
	if len(d.resultsRaw) == 0 {
		return nil
	}
	for _, item := range d.resultsRaw {
		res = FilmIDRegexp.FindStringSubmatch(item)
		if len(res) > 0 {
			id = res[1]
			i, err := strconv.Atoi(id)
			if err != nil {
				continue
			}
			d.FilmIDs = append(d.FilmIDs, i)
			continue
		}
		res = PersonIDRegexpt.FindStringSubmatch(item)
		if len(res) > 0 {
			id = res[1]
			i, err := strconv.Atoi(id)
			if err != nil {
				continue
			}
			d.PersonIDs = append(d.PersonIDs, i)
			continue
		}
	}
	return nil
}

func (d *Data) processObjects() error {
	if len(d.objectsRaw) == 0 {
		return nil
	}
	for _, item := range d.objectsRaw {
		if strings.Index(item, "COBJECT") >= 0 {
			f := Film{}
			if err := json.Unmarshal([]byte(item), &f); err != nil {
				return err
			}
			f.URL = fmt.Sprintf(FilmURL, f.EntityId)
			d.Films = append(d.Films, f)
		} else if strings.Index(item, "PERSON") >= 0 {
			p := Person{}
			if err := json.Unmarshal([]byte(item), &p); err != nil {
				return err
			}
			p.URL = fmt.Sprintf(PersonURL, p.EntityId)
			d.Persons = append(d.Persons, p)
		} else {
			return errors.New("Unknown search object type")
		}
	}
	return nil
}

func (i *Image) Size(size string) string {
	return fmt.Sprintf("%s/%s", i.BaseUrl, size)
}

func Query(query string) (*Data, error) {
	ddgUrl := encodeUrl(query)
	body, err := do(ddgUrl)
	if err != nil {
		return nil, err
	}
	d := &Data{}
	if err = d.decode(body); err != nil {
		return nil, err
	}
	if err = d.processIDs(); err != nil {
		return nil, err
	}
	if err = d.processObjects(); err != nil {
		return nil, err
	}
	return d, nil
}
