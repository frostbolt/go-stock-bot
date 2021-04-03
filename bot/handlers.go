package bot

import (
	"fmt"
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

func inlineQueryHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	normalizedTicker := strings.TrimPrefix(update.InlineQuery.Query, "$")
	log.Printf("Inline Query handler, %s", normalizedTicker)

	if !SearchQueryRegex.Match([]byte(normalizedTicker)) {
		return
	}

	searchResults, err := apisdk.SearchQuotes(normalizedTicker)
	if err != nil {
		log.Error(err)
		return
	}

	inlineQueryResults := make([]interface{}, 0, len(searchResults.Quotes))

	for key, quote := range searchResults.Quotes {
		var articleName string
		if len(quote.Name) > 0 {
			articleName = quote.Name
		} else if len(quote.Shortname) > 0 {
			articleName = quote.Shortname
		} else if len(quote.Longname) > 0 {
			articleName = quote.Longname
		} else {
			articleName = quote.Symbol
		}

		article := tgbotapi.NewInlineQueryResultArticle(
			fmt.Sprint(key),
			articleName,
			quote.Symbol,
		)
		article.Description = quote.Symbol
		article.InputMessageContent = tgbotapi.InputTextMessageContent{
			Text: fmt.Sprintf("$%s", quote.Symbol),
		}

		inlineQueryResults = append(inlineQueryResults, article)
	}

	response := tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		Results:       inlineQueryResults,
	}

	if _, err = bot.AnswerInlineQuery(response); err != nil {
		log.Error(err)
	}
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
