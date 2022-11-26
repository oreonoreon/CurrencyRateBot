package bot

import (
	"CurrencyRateBot/internal/start"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"reflect"
	"regexp"
)

func LOGIKA(ID int64, message string) error {
	item, found := botCache.Get(ID)
	if !found {
		msg := tgbotapi.NewMessage(ID, "Напишите сумму которую желаете поменять.")
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
		s := NewStatement()
		s.amountSet = true
		botCache.Set(ID, s, 0)

	} else {
		if s, ok := item.(*botStatement); !ok {
			err := fmt.Errorf("%v can't be asserted to %v", reflect.TypeOf(item), reflect.TypeOf(botStatement{}))
			return err
		} else {
			err := s.Cases(ID, message)
			return err
		}
	}
	return nil
}

func regularExpCheck(str string) (string, error) {
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		return "", err
	}
	return re.FindString(str), nil
}

type botStatement struct {
	bank      string
	amount    string
	amountSet bool
	bankSet   bool
}

func NewStatement() *botStatement {
	return &botStatement{}
}

func (s *botStatement) Cases(ID int64, message string) error {

	switch {
	case !s.amountSet && !s.bankSet:
		msg := tgbotapi.NewMessage(ID, "Напишите сумму которую желаете поменять.")
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
		s.amountSet = true
	case s.amountSet && !s.bankSet:
		amount, err := regularExpCheck(message)
		if err != nil {
			s.amountSet = false
			return err
		}
		s.amount = amount
		msg := tgbotapi.NewMessage(ID, "Выбирите банк для перевода.")
		msg.ReplyMarkup = replyKbBank
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
		s.bankSet = true
	case s.amountSet && s.bankSet:
		if message == bank || message == sberbank {
			s.bank = message
		} else {
			s.bankSet = false
			msg := tgbotapi.NewMessage(ID, "Параметры запроса введены не верно. Попробуйте ещё раз")
			bot.Send(msg)
			return fmt.Errorf("параметры запроса введены не верно - %v", message)
		}
		rate, err := start.Start(s.bank, s.amount)
		if err != nil {
			msg := tgbotapi.NewMessage(ID, err.Error())
			bot.Send(msg)
			s.amountSet = false
			return err
		}

		s.amountSet = false
		s.bankSet = false

		msg := tgbotapi.NewMessage(ID, fmt.Sprint(rate))
		msg.ReplyMarkup = replyKeyboard
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
