package main

import (
	"encoding/json"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/leominov/gokinopoisk/search"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
}

var theatres  = []string{
	"Космос",
	"Родина",
	"Октябрь",
	"Ветразь",
	"Чырвоная Зорка",

}


type FullMoviesInfo struct {
	search.Film
	theaters map[string]string
	imbdRate string
}

func main() {
	//open config file and read data to Config struct var
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err := decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	//start new Telegram Bot with API token from Config struct var
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		typing := tgbotapi.NewChatAction(update.Message.Chat.ID, "typing")
		bot.Send(typing)
		if update.Message.Command() == "go" {
			var msgtext string
			movies := getMovies()
			getMoviesData(*movies)
			fullInfo := parseMoviePage(*movies)



			for _, movie := range fullInfo {
				rating := fmt.Sprint(movie.Rating.Rate)
				if rating == "0" {
					rating = movie.imbdRate
				}
				msgtextPart := ""
				for _, theatre := range theatres{
					data, exist := movie.theaters[theatre]
					if exist{
						msgtextPart += theatre + data + ", "


					}
				}

				msgtext += "<a href='" + movie.URL + "'>" + movie.Title + "</a>  <b>" + rating + "</b>\n <code>   "+ msgtextPart + "</code>\n"

			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtext)
			msg.ParseMode = "HTML"
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		if update.Message.Command() == "start" {
			msgtxt := "Привет. Отправь мне команду /go и получи список фильмов в прокате города Могилева. "
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtxt)
			bot.Send(msg)
		}
		if update.Message.Command() == "help" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Помоги себе сам")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		//bot.Send(msg)
	}

}
