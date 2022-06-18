package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	helpText := `/help - help
	/list - list products
	`
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpText)
	c.bot.Send(msg)
}
