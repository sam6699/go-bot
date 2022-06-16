package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sam6699/bot/internal/service/product"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{Timeout: 60}
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(productService, bot, update.Message)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	helpText := `/help - help
	/list - list products
	`

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpText)
	bot.Send(msg)
}

func listCommand(productService *product.Service, bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	products := productService.List()
	msgText := "Products: \n\n"
	for _, p := range products {
		msgText += p.Title
		msgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Your message "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID
	bot.Send(msg)
}
