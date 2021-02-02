package bot

import (
	"regexp"
	"strings"

	APISDK "github.com/frostbolt/go-stock-bot/api-sdk"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	log "github.com/sirupsen/logrus"
)

// Command describes the functions that are used for incoming messages handling
type Command func(*tgbotapi.BotAPI, tgbotapi.Update)

// Handlers is a map of custom bot commands
var handlers = map[string]Command{
	"/start": start,
	"/info":  info,
}

func defaultHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Default handler")
	tickerRegex, _ := regexp.Compile("^\\$[A-Z\\-]{1,7}$")

	if update.Message.Text[0] != '$' && !tickerRegex.Match([]byte(update.Message.Text)) {
		return
	}

	tickerInfo, err := APISDK.GetInfoOnTicker(strings.TrimPrefix(update.Message.Text, "$"))
	var response string
	if err != nil {
		response = "¯\\_(ツ)_/¯"
	} else {
		response = FormatTickerResults(tickerInfo)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}

// GetHandler returns a propper handler
func GetHandler(command string) Command {
	function, isDefined := handlers[command]
	if !isDefined {
		function = defaultHandler
	}

	return function
}

func start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Start command")
}

func info(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Info command")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Author: @frostbolt")
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
