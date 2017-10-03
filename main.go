package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func main()  {
	bot, err := tgbotapi.NewBotAPI("437757616:AAErig4Hb9ZZoVhS5CnTUPaI4DbsCLl5Q3E")
	if err != nil{
		log.Fatal(err)
	}

	bot.Debug = true
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
			var msgstr string
			for _, movie := range getMovies() {
				msgstr += movie.name + ":   " + movie.rating +"\n"

			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgstr)
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
