package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"time"
)

var botCache *Cache
var bot *tgbotapi.BotAPI

func init() {
	//сосдадим кэш бота
	botCache = New(10*time.Minute, 60*time.Minute)

	fmt.Println("Bot cached initialised")
	const GetYourCurrencyRate string = "5655548481:AAFlSdYiyhgf7VX1jX5k5h2WsA6Fu-RIUZI"
	os.Setenv("token", GetYourCurrencyRate)
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

}

func newBot() tgbotapi.UpdatesChannel {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return updates
}
func Bot() {
	//получаем апдейты
	updates := newBot()
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				msg := tgbotapi.NewMessage(update.SentFrom().ID, "Обмен Рублей на Таиский бат")
				msg.ReplyMarkup = replyKeyboard
				bot.Send(msg)
			} else {
				err := LOGIKA(update.SentFrom().ID, update.Message.Text)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
