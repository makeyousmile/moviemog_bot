package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	)

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func getMovies()  *[]string{

	movies := make([]string,0)

	doc, err := goquery.NewDocument("https://afisha.tut.by/film-mogilev/")
	if err != nil {
		log.Fatal(err)
	}


	selection := doc.Find("div#events-block").First()
		selection.Find("a.name").Each(func(i int, selection *goquery.Selection) {
			movies = append(movies, selection.Text())
		})

	reverse(movies)

	return &movies
}
