package start

import (
	"CurrencyRateBot/internal/OuterAPI/BinanceP2P"
	"errors"
	"math"
	"strconv"
)

func Start(bank, amount string) (float64, error) {

	pb := BinanceP2P.NewPostBody("RUB", []string{bank}, "BUY", amount)
	res1, _ := BinanceP2P.Request(pb)
	pb2 := BinanceP2P.NewPostBody("THB", []string{"BANK"}, "SELL", "")
	res2, _ := BinanceP2P.Request(pb2)
	if len(res1) == 0 || len(res2) == 0 {
		return 0, errors.New("На ваш запрос не найдено возможных курсов обмена, попробуйте изменить пораметры запроса.")
	}
	price1, err := strconv.ParseFloat(res1[0].Adv.Price, 64)
	if err != nil {
		return 0, err
	}
	price2, err := strconv.ParseFloat(res2[0].Adv.Price, 64)
	if err != nil {
		return 0, err
	}
	rate := price1 / price2
	//customerRate := math.Ceil((rate*1.02)*100) / 100
	customerRate := math.Round((rate*1.02)*100) / 100
	return customerRate, nil
}
