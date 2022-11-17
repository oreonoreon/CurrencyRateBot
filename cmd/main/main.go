package main

import (
	"CurrencyRateBot/internal/OuterAPI/BinanceP2P"
	"fmt"
	"time"
)

func main() {

	pb := BinanceP2P.NewPostBody("RUB", []string{"TinkoffNew"}, "BUY", "")
	res1, _ := BinanceP2P.Request(pb)
	pb2 := BinanceP2P.NewPostBody("THB", []string{"BANK"}, "SELL", "")
	res2, _ := BinanceP2P.Request(pb2)
	//price1, _ := strconv.ParseFloat(res1[0].Adv.Price, 64)
	//price2, _ := strconv.ParseFloat(res2[0].Adv.Price, 64)
	//rate := math.Ceil(price1/price2*100) / 100
	for _, v := range res1 {
		fmt.Println(v.Adv.Price)
	}
	fmt.Println("-----------------THB--------------")
	for _, v := range res2 {
		fmt.Println(v.Adv.Price)
	}
	time.Sleep(5 * time.Second)

	//for {
	//	BinanceP2P.Request()
	//	in := bufio.NewReader(os.Stdin)
	//	var restart string
	//	fmt.Fscan(in, &restart)
	//	if restart == "r" || restart == "R" {
	//		continue
	//	} else {
	//		break
	//	}
	//
	//}

}
