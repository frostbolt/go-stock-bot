package main

import (
	"github.com/frostbolt/go-stock-bot/bot"
	conf "github.com/frostbolt/go-stock-bot/configure"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Debug("Starting the application")
	conf.LoadConfig()
	config := conf.GetConfig()

	bot.RunBot(config.GetString("bot.token"), config.GetInt("bot.timeout"), config.GetBool("bot.debug"))
}
