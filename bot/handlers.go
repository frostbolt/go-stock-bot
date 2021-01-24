package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

// Command describes the functions that are used for incoming messages handling
type Command func(*tgbotapi.BotAPI, tgbotapi.Update)

// Handlers is a map of custom bot commands
var handlers = map[string]Command{
	"/start": start,
}

func defaultHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("Default handler")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
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
