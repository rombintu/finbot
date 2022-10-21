package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/rombintu/finbot/internal"
	"github.com/rombintu/finbot/tools"
)

const (
	RequestSuccess string = "Операция прошла успешно"
	RequestError   string = "Неправильный ввод. Стоимость Категория [Комментарий]"
	DatabaseError  string = "500 internal error"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	token := os.Getenv("TOKEN")
	if uri == "" || token == "" {
		log.Fatal("You must set your 'MONGODB_URI' and 'TOKEN' environmental variable")
	}

	store := internal.NewStore(uri)

	bot, err := tgbotapi.NewBotAPI(token)
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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			note, ok := tools.Filter(update.Message.Text)
			if ok {
				note.UUID = update.Message.Chat.ID
				if err := store.PutNote(note); err != nil {
					msg.Text = DatabaseError
					continue
				}
				msg.Text = RequestSuccess
			} else {
				msg.Text = RequestError
			}
			// msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
