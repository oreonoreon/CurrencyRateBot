package main

import (
	"CurrencyRateBot/internal/bot"
	"CurrencyRateBot/internal/logger"
)

func init() {
	logger.Debug = true
	logger.Init()
	bot.WebhookOn = true
	bot.Init()
}
func main() {
	bot.Bot()
}
