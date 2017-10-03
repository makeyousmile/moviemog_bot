package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func getMovies() string {
	doc, err := goquery.NewDocument("http://afisha.tut.by/film-mogilev/")
	if err != nil {
		log.Fatal(err)
	}
	var out string = "aloha"
	// Find the review items
	doc.Find(".events-block js-cut_wrapper").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
		out = band
	})
	return out
}