package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func newBot() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	os.Setenv("token", halyava007BotToken)
	bot, err := tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return bot, updates
}
func Bot() {
	//конектимся к телеграм и получаем апдейты
	newBot()

}
