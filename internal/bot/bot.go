package bot

import (
	"CurrencyRateBot/internal/logger"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
	"time"
)

const webhook string = "https://effulgent-pegasus-fb77da.netlify.app/"

var WebhookOn bool
var port string

var botCache *Cache
var bot *tgbotapi.BotAPI

func Init() {
	//получим порт из переменой окружения
	port = os.Getenv("PORT")
	fmt.Printf("PORT SET TO -> %v\n", port)
	//создадим кэш бота
	botCache = New(10*time.Minute, 60*time.Minute)
	fmt.Println("Bot cached initialised")

	const GetYourCurrencyRate string = "5655548481:AAFlSdYiyhgf7VX1jX5k5h2WsA6Fu-RIUZI"
	os.Setenv("token", GetYourCurrencyRate)

	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		logger.Log.Panic(err)
	}
	bot.Debug = false
	logger.Log.Printf("Authorized on account %s", bot.Self.UserName)
}

func newBot() tgbotapi.UpdatesChannel {
	if WebhookOn {
		wh, err := tgbotapi.NewWebhook(webhook)
		if err != nil {
			logger.Log.Fatal(err)
		}

		if _, err = bot.Request(wh); err != nil {
			logger.Log.Fatal(err)
		}
		info, err := bot.GetWebhookInfo()
		if err != nil {
			logger.Log.Fatal(err)
		}

		if info.LastErrorDate != 0 {
			logger.Log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
		}

		updates := bot.ListenForWebhook("/")
		go func() {
			if err := http.ListenAndServe(":"+port, nil); err != nil {
				log.Fatal(err)
			}
		}()
		return updates
	} else {
		wh := tgbotapi.DeleteWebhookConfig{DropPendingUpdates: true}
		if _, err := bot.Request(wh); err != nil {
			logger.Log.Fatal(err)
		}
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)
		return updates
	}

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
					logger.Log.Println(err)
				}
			}
		}
	}
}
