package main

import (
	"github.com/leominov/gokinopoisk/search"
	"log"
	"sort"
)

//func getMoviesData(movieTittles []search.Film) *[]search.Film {
//
//	movies := make([]search.Film,0)
//	log.Print(movieTittles)
// 	for _, movieTittle := range movieTittles{
//		res, err := search.Query(movieTittle)
//		if err != nil { log.Panic(err) }
//
//		//Выбор нужного фильма (более новый)
//		for _, film := range res.Films {
//
//			if len(film.Years) > 0 && film.Type == "MOVIE" {
//				if film.Years[0] > 2015{
//
//					movies = append(movies, film)
//				}
//
//			}
//		}
//	}
//
//return &movies
//}
func getMoviesData(movies []search.Film) {
	for i, movie := range movies{
		res, err := search.Query(movie.Title)
		if err != nil{log.Fatal(err)}
		for _, film := range res.Films{
			if len(film.Years) > 0 && film.Type == "MOVIE" {
				if film.Years[0] > 2015{
					movies[i] = film
					}

				}
		}
	}
	sort.Slice(movies, func(i, j int) bool {
		switch movies[i].Rating.Rate > movies[j].Rating.Rate{
		case true:
			return true
		case false:
			return false
		}
		return movies[i].Rating.Rate > movies[j].Rating.Rate
	})

}