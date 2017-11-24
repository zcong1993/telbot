package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	bfx "github.com/zcong1993/telbot/bitfinex"
	"regexp"
)

func main() {
	symbols := []string{bitfinex.BTCUSD, "BCHUSD", bitfinex.ETHUSD, bitfinex.LTCUSD, bitfinex.ETCUSD, "BCHBTC", bitfinex.ETCBTC, bitfinex.ETHBTC}
	b := bfx.NewBfx(symbols)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Println(err.Error())
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		echo := regexp.MustCompile(`\/query`)

		result := echo.MatchString(update.Message.Text)

		if result {
			prices, err := b.GetTicker()
			var text string
			if err != nil {
				text = "An error occurred"
				log.Println(err.Error())
			} else {
				d := [][]string{}
				for k, v := range prices {
					d = append(d, []string{k, fmt.Sprintf("%f", v)})
				}
				buf := CreateTableText(d)
				text = buf.String()
				fmt.Println(text)
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}
