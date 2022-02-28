package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("5256302597:AAHspvcoZKgQAbH6WsiqWZDwj5UFuszPqfA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			text := "Greetings, " + update.Message.Chat.FirstName

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

			bot.Send(msg)
		}
	}

}