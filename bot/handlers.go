package bot

import (
	"strings"

	apisdk "github.com/frostbolt/go-stock-bot/apisdk"
	"github.com/frostbolt/go-stock-bot/formatters"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	log "github.com/sirupsen/logrus"
)

// Command describes the functions that are used for incoming messages handling
type Command func(*tgbotapi.BotAPI, tgbotapi.Update)

// Handlers is a map of custom bot commands
var Handlers = map[string]Command{
	"start": start,
	"info":  info,
}

func defaultHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	normalizedTicker := strings.TrimPrefix(update.Message.Text, "$")
	log.Printf("Default handler, %s", normalizedTicker)

	tickerInfo, err := apisdk.GetInfoOnTicker(normalizedTicker)

	if err != nil {
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, formatters.FormatTickerResults(tickerInfo))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

func start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Start command")
}

func info(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Info command")
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Author: @frostbolt \nRepo: https://github.com/frostbolt/go-stock-bot",
	)

	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
