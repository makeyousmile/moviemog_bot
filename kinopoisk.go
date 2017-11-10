package main

import (
	"github.com/leominov/gokinopoisk/search"
	"log"
)

func getRating(movieTittle string) float32 {
	var rating float32
	res, err := search.Query(movieTittle)
	if err != nil {
		log.Panic(err)
	}
	//Выбор нужного фильма (более новый)
	maxYear := 0
	rating = 0
	for _, film := range res.Films {
		if film.Years[0] >  maxYear{
			rating = film.Rating.Rate
			maxYear = film.Years[0]
		}
	}

return rating
}