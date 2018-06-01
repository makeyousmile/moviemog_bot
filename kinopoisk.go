package main

import (
	"github.com/leominov/gokinopoisk/search"
	"log"
)

func getMoviesData(movieTittles []string) *[]search.Film {

	movies := make([]search.Film,0)
	log.Print(movieTittles)
 	for _, movieTittle := range movieTittles{
		res, err := search.Query(movieTittle)
		if err != nil { log.Panic(err) }

		//Выбор нужного фильма (более новый)
		for _, film := range res.Films {

			if len(film.Years) > 0 && film.Type == "MOVIE" {
				if film.Years[0] > 2015{

					movies = append(movies, film)
				}

			}
		}
	}

return &movies
}