package main

import (
	"github.com/frostbolt/go-stock-bot/bot"
	"github.com/frostbolt/go-stock-bot/cache"
	conf "github.com/frostbolt/go-stock-bot/configure"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Debug("Starting the application")
	conf.LoadConfig()
	config := conf.GetConfig()

	cache.SetupDBConnection(
		config.GetString("redis.addr"),
		config.GetString("redis.password"),
		config.GetInt("redis.DB"),
		config.GetInt64("redis.expiration"),
	)

	bot.RunBot(
		config.GetString("bot.token"),
		config.GetInt("bot.timeout"),
		config.GetBool("bot.debug"),
	)
}
