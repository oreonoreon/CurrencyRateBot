package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var replyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Какой курс?"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Сообщить о проблеме."),
	),
)
