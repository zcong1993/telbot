package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
	"regexp"
)

func main() {
	symbols := []string{"USDT_REP", "USDT_BCH", "BTC_ETC", "USDT_ETH", "USDT_LTC", "USDT_BTC", "BTC_ETH", "USDT_ETC", "BTC_BCH"}
	polo := NewPolo(symbols)

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
			prices, err := polo.GetPrices()
			var text string
			if err != nil {
				text = "An error occurred"
				log.Println(err.Error())
			} else {
				buf := CreateTableText(prices)
				text = buf.String()
				fmt.Println(text)
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}
