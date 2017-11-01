package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"encoding/json"
	"os"
	)

type Config struct {
	TelegramBotToken string
}

func main()  {
	//token := json.Decoder{os.Open("config.json")}
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err := decoder.Decode(&cfg)
	if err !=nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil{
		log.Fatal(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)

	for update := range updates{
		if update.Message == nil{
			continue
		}

		log.Printf("[%s]%s", update.Message.From.FirstName,update.Message.Command())


		if update.Message.Command() == "go"{
			var msgtext string
			for _, movie := range *getMovies() {
				msgtext += movie.name + ":   " + movie.rating +"\n"

			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtext)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		if update.Message.Command() == "help"{
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Помоги себе сам")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		//bot.Send(msg)
	}

}
