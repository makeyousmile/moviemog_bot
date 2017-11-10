package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

type movie struct {
	name string
	rating string
	kinopoiskRating string
	time string
	price string
}

func getMovies()  *[]movie{

	movies := make([]movie,0)

	doc, err := goquery.NewDocument("https://afisha.tut.by/film-mogilev/")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li.lists__li").Each(func(i int, selection *goquery.Selection) {
		var m movie

		val, exist := selection.Attr("itemtype")
		if exist && val == "http://data-vocabulary.org/Event"{

			selection.Find("span").Each(func(n int, selection *goquery.Selection) {
				val, exist := selection.Attr("itemprop")
				if exist && val == "summary"{
					m.name = selection.Text()
					m.kinopoiskRating = fmt.Sprint(getRating(m.name))
					m.rating = selection.Parent().Parent().Find("span.raiting").Text()
					movies = append(movies, m)
				}
			})

			//selection.Find("a.media span").Each(func(count int, selection *goquery.Selection) {
			//	log.Print(selection.Html())
			//	log.Print(count)
			//})
		}


	})

	return &movies
}
