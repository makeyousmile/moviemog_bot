package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"github.com/leominov/gokinopoisk/search"
)


func getMovies()  *[]search.Film{

	movies := make([]search.Film,0)

	doc, err := goquery.NewDocument("https://afisha.tut.by/film-mogilev/")
	if err != nil {
		log.Fatal(err)
	}


	selection := doc.Find("div#events-block").First()
		selection.Find("a.name").Each(func(i int, selection *goquery.Selection) {
			var movie search.Film
			movie.Title = selection.Text()
			val, exist := selection.Attr("href")
			if exist{
				movie.URL = val
			}
			movies = append(movies, movie)
		})



	return &movies
}
