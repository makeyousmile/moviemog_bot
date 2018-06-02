package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/leominov/gokinopoisk/search"
	"log"
	"strings"
)

func getMovies() *[]search.Film {

	movies := make([]search.Film, 0)

	doc, err := goquery.NewDocument("https://afisha.tut.by/film-mogilev/")
	if err != nil {
		log.Panic(err)
	}

	selection := doc.Find("div#events-block").First()
	selection.Find("a.name").Each(func(i int, selection *goquery.Selection) {
		var movie search.Film
		movie.Title = selection.Text()
		val, exist := selection.Attr("href")
		if exist {
			movie.URL = val
		}
		movies = append(movies, movie)
	})

	return &movies
}
func parseMoviePage(movies []search.Film) [100]FullMoviesInfo {
	var fullInfo [100]FullMoviesInfo

	for i, movie := range movies {
		fullInfo[i].Film = movie
		fullInfo[i].theaters = make(map[string]string, 6)
		doc, err := goquery.NewDocument(movie.URL)

		if err != nil {
			log.Panic(err)
		}
		imbdRate := doc.Find("td.IMDb").Find("b").First().Text()

		fullInfo[i].imbdRate = imbdRate

		selection := doc.Find("div.b-film-info").First()
		selection.Find("li.b-film-list__li").Each(func(_ int, selection *goquery.Selection) {
			theater := strings.TrimSpace(selection.Find("div.film-name").Text())
			selection.Find("a.tooltip-holder").Each(func(_ int, selection *goquery.Selection) {
				hour, exist := selection.Parent().Attr("data-hour")
				if exist {

					fullInfo[i].theaters[theater] += " "+ hour +":"
				}
				minute, exist := selection.Parent().Attr("data-minute")
				if exist {
					if minute == "0"{
						minute = "00"
					}
					fullInfo[i].theaters[theater] += minute
				}
			})


		})


	}

	return fullInfo
}
