package main

import (
	"CurrencyRateBot/internal/bot"
	"CurrencyRateBot/internal/logger"
)

func init() {
	logger.Debug = true
}
func main() {
	bot.Bot()
}
