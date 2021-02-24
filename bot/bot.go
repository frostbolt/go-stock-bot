package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var bot *tgbotapi.BotAPI

// GetBotName returns the bot name
func GetBotName() string {
	return bot.Self.UserName
}

// RunBot launches the bot updates listener
func RunBot(token string, timeout int, debug bool) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = timeout

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		var handler = GetHandler(update.Message.Text)
		handler(bot, update)
	}
}
