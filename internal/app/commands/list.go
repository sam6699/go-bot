package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	products := c.productService.List()
	msgText := "Products: \n\n"
	for _, p := range products {
		msgText += p.Title
		msgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	c.bot.Send(msg)
}
