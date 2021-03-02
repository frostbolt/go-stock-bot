package bot

import (
	"fmt"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var bot *tgbotapi.BotAPI
var botName string

// GetBotName returns the bot username
func GetBotName() string {
	return botName
}

// BotCommandRegex is meant to match bot commands, returns a group, that contains the command and the bot it relates to
var BotCommandRegex *regexp.Regexp

// TickerRegex is meant to match the ticker
var TickerRegex *regexp.Regexp

// InitializeBotVariables gets the bot instance and setups the propper environment
func InitializeBotVariables(bot *tgbotapi.BotAPI) {
	botName = fmt.Sprintf("@%s", bot.Self.UserName)
	BotCommandRegex = regexp.MustCompile(fmt.Sprintf("^/(?P<command>[a-zA-Z]+)(?P<botname>%s)?$", botName))
	TickerRegex = regexp.MustCompile("^\\$[A-Z\\-]{1,7}$")
}

// CommandValidator returns validated command and the flag that indicates whether the validation was successful
func CommandValidator(update *tgbotapi.Update) (string, bool) {
	if update.Message == nil || update.Message.Text == "" {
		return "", false
	}

	submatch := BotCommandRegex.FindStringSubmatch(update.Message.Text)
	match := map[string]string{}
	for k, v := range submatch {
		match[BotCommandRegex.SubexpNames()[k]] = v
	}

	if update.Message.Text[0] == '$' && !TickerRegex.Match([]byte(update.Message.Text)) {
		return "defaultHandler", true
	}

	command, isCommandDefined := match["command"]
	botUserName, _ := match["botname"]
	log.Println(match, GetBotName(), botUserName == GetBotName())

	if isCommandDefined && (update.Message.Chat.Type == "private" || botUserName == GetBotName()) {
		return command, true
	}

	return "", false
}

// GetHandler returns a propper handler and the flag which shows whether the handler is correct
func GetHandler(update *tgbotapi.Update) (Command, bool) {
	command, isValidationSuccessful := CommandValidator(update)

	function, isDefined := Handlers[command]
	if !isDefined {
		function = defaultHandler
	}

	return function, isValidationSuccessful
}

// RunBot launches the bot updates listener
func RunBot(token string, timeout int, debug bool) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	InitializeBotVariables(bot)
	log.Printf("Authorized on account %s", botName)

	bot.Debug = debug

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = timeout

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		var handler, isHandlerDefined = GetHandler(&update)
		if !isHandlerDefined {
			continue
		}

		handler(bot, update)
	}
}
