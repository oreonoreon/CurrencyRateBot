package BinanceP2P

import (
	"bytes"
	"encoding/json"
)

const (
	host   = "p2p.binance.com"
	method = "bapi/c2c/v2/friendly/c2c/adv/search"
)

/*
PostBody example of JSON which should be and description of fields
POST Fields
{
"asset": "USDT",
"fiat": "THB",
"merchantCheck": true, //false
"page": 1,
"payTypes": ["BANK"],
"publisherType": "merchant", // null
"rows": 20,
"tradeType": "SELL",
"transAmount":  "5000"
}

asset: Currently available asset USDT, BTC, BNB, BUSD, ETH, DAI.

fiat: Its long, visit sanchezmarcos.

merchantCheck: Well i don't know its use, but value is null, true, false.

page: The endpoint is paginated.

payTypes: An array of payment types example BANK, GoMoney, CashDeposit etc.

payTypes depend on fiat used so you might not see some of these but there's a lot of payment types.

publisherType: I'm only aware of merchant.

row: Amount of rows from 1 - 20.

tradeType: BUY or SELL.

transAmount: Filter merchant by amount.
*/
type PostBody struct {
	Asset         string   `json:"asset"`
	Fiat          string   `json:"fiat"`
	MerchantCheck bool     `json:"merchantCheck"`
	Page          int      `json:"page"`
	PayTypes      []string `json:"payTypes"`
	PublisherType string   `json:"publisherType,omitempty"`
	Rows          int      `json:"rows"`
	TradeType     string   `json:"tradeType"`
	TransAmount   string   `json:"transAmount,omitempty"`
}

func NewPostBody(

	fiat string,

	payTypes []string,

	tradeType string,
	transAmount string,
) PostBody {
	return PostBody{
		Asset:         "USDT",
		Fiat:          fiat,
		MerchantCheck: false,
		Page:          1,
		PayTypes:      payTypes,
		PublisherType: "",
		Rows:          5,
		TradeType:     tradeType,
		TransAmount:   transAmount,
	}
}
func (pb PostBody) Encode() *bytes.Buffer {
	Buf := new(bytes.Buffer)
	json.NewEncoder(Buf).Encode(pb)

	return Buf
}

type ResBody struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          []struct {
		Adv struct {
			AdvNo                 string      `json:"advNo"`
			Classify              string      `json:"classify"`
			TradeType             string      `json:"tradeType"`
			Asset                 string      `json:"asset"`
			FiatUnit              string      `json:"fiatUnit"`
			AdvStatus             interface{} `json:"advStatus"`
			PriceType             interface{} `json:"priceType"`
			PriceFloatingRatio    interface{} `json:"priceFloatingRatio"`
			RateFloatingRatio     interface{} `json:"rateFloatingRatio"`
			CurrencyRate          interface{} `json:"currencyRate"`
			Price                 string      `json:"price"`
			InitAmount            interface{} `json:"initAmount"`
			SurplusAmount         string      `json:"surplusAmount"`
			AmountAfterEditing    interface{} `json:"amountAfterEditing"`
			MaxSingleTransAmount  string      `json:"maxSingleTransAmount"`
			MinSingleTransAmount  string      `json:"minSingleTransAmount"`
			BuyerKycLimit         interface{} `json:"buyerKycLimit"`
			BuyerRegDaysLimit     interface{} `json:"buyerRegDaysLimit"`
			BuyerBtcPositionLimit interface{} `json:"buyerBtcPositionLimit"`
			Remarks               interface{} `json:"remarks"`
			AutoReplyMsg          string      `json:"autoReplyMsg"`
			PayTimeLimit          interface{} `json:"payTimeLimit"`
			TradeMethods          []struct {
				PayId                interface{} `json:"payId"`
				PayMethodId          string      `json:"payMethodId"`
				PayType              interface{} `json:"payType"`
				PayAccount           interface{} `json:"payAccount"`
				PayBank              interface{} `json:"payBank"`
				PaySubBank           interface{} `json:"paySubBank"`
				Identifier           string      `json:"identifier"`
				IconUrlColor         interface{} `json:"iconUrlColor"`
				TradeMethodName      string      `json:"tradeMethodName"`
				TradeMethodShortName *string     `json:"tradeMethodShortName"`
				TradeMethodBgColor   string      `json:"tradeMethodBgColor"`
			} `json:"tradeMethods"`
			UserTradeCountFilterTime        interface{}   `json:"userTradeCountFilterTime"`
			UserBuyTradeCountMin            interface{}   `json:"userBuyTradeCountMin"`
			UserBuyTradeCountMax            interface{}   `json:"userBuyTradeCountMax"`
			UserSellTradeCountMin           interface{}   `json:"userSellTradeCountMin"`
			UserSellTradeCountMax           interface{}   `json:"userSellTradeCountMax"`
			UserAllTradeCountMin            interface{}   `json:"userAllTradeCountMin"`
			UserAllTradeCountMax            interface{}   `json:"userAllTradeCountMax"`
			UserTradeCompleteRateFilterTime interface{}   `json:"userTradeCompleteRateFilterTime"`
			UserTradeCompleteCountMin       interface{}   `json:"userTradeCompleteCountMin"`
			UserTradeCompleteRateMin        interface{}   `json:"userTradeCompleteRateMin"`
			UserTradeVolumeFilterTime       interface{}   `json:"userTradeVolumeFilterTime"`
			UserTradeType                   interface{}   `json:"userTradeType"`
			UserTradeVolumeMin              interface{}   `json:"userTradeVolumeMin"`
			UserTradeVolumeMax              interface{}   `json:"userTradeVolumeMax"`
			UserTradeVolumeAsset            interface{}   `json:"userTradeVolumeAsset"`
			CreateTime                      interface{}   `json:"createTime"`
			AdvUpdateTime                   interface{}   `json:"advUpdateTime"`
			FiatVo                          interface{}   `json:"fiatVo"`
			AssetVo                         interface{}   `json:"assetVo"`
			AdvVisibleRet                   interface{}   `json:"advVisibleRet"`
			AssetLogo                       interface{}   `json:"assetLogo"`
			AssetScale                      int           `json:"assetScale"`
			FiatScale                       int           `json:"fiatScale"`
			PriceScale                      int           `json:"priceScale"`
			FiatSymbol                      string        `json:"fiatSymbol"`
			IsTradable                      bool          `json:"isTradable"`
			DynamicMaxSingleTransAmount     string        `json:"dynamicMaxSingleTransAmount"`
			MinSingleTransQuantity          string        `json:"minSingleTransQuantity"`
			MaxSingleTransQuantity          string        `json:"maxSingleTransQuantity"`
			DynamicMaxSingleTransQuantity   string        `json:"dynamicMaxSingleTransQuantity"`
			TradableQuantity                string        `json:"tradableQuantity"`
			CommissionRate                  string        `json:"commissionRate"`
			TradeMethodCommissionRates      []interface{} `json:"tradeMethodCommissionRates"`
			LaunchCountry                   interface{}   `json:"launchCountry"`
			AbnormalStatusList              interface{}   `json:"abnormalStatusList"`
			CloseReason                     interface{}   `json:"closeReason"`
		} `json:"adv"`
		Advertiser struct {
			UserNo           string        `json:"userNo"`
			RealName         interface{}   `json:"realName"`
			NickName         string        `json:"nickName"`
			Margin           interface{}   `json:"margin"`
			MarginUnit       interface{}   `json:"marginUnit"`
			OrderCount       interface{}   `json:"orderCount"`
			MonthOrderCount  int           `json:"monthOrderCount"`
			MonthFinishRate  float64       `json:"monthFinishRate"`
			AdvConfirmTime   interface{}   `json:"advConfirmTime"`
			Email            interface{}   `json:"email"`
			RegistrationTime interface{}   `json:"registrationTime"`
			Mobile           interface{}   `json:"mobile"`
			UserType         string        `json:"userType"`
			TagIconUrls      []interface{} `json:"tagIconUrls"`
			UserGrade        int           `json:"userGrade"`
			UserIdentity     string        `json:"userIdentity"`
			ProMerchant      interface{}   `json:"proMerchant"`
			IsBlocked        interface{}   `json:"isBlocked"`
		} `json:"advertiser"`
	} `json:"data"`
	Total   int  `json:"total"`
	Success bool `json:"success"`
}
