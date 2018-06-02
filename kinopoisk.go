package main

import (
	"github.com/leominov/gokinopoisk/search"
	"log"
)


func getMoviesData(movies []search.Film) {
	for i, movie := range movies {
		res, err := search.Query(movie.Title)
		if err != nil {
			log.Fatal(err)
		}
		for _, film := range res.Films {
			if len(film.Years) > 0 && film.Type == "MOVIE" {
				if film.Years[0] > 2015 {
					tempUrl := movies[i].URL
					movies[i] = film
					movies[i].URL = tempUrl
				}

			}
		}
	}

}
