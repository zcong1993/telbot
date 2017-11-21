package main

import (
	"github.com/jyap808/go-poloniex"
)

const (
	// API_KEY is poloniex api key
	API_KEY = ""
	// API_SECRET is poloniex api secret
	API_SECRET = ""
)

// Polo is goi poloniex client
type Polo struct {
	Symbols          []string
	AvailableSymbols []string
	tickers          map[string]poloniex.Ticker
	poloniex         *poloniex.Poloniex
}

// NewPolo is Polo client constructor
func NewPolo(symbols []string) *Polo {
	poloniex := poloniex.New(API_KEY, API_SECRET)
	polo := &Polo{Symbols: symbols, poloniex: poloniex}
	return polo
}

// GetPrices return filtered price results
func (polo *Polo) GetPrices() ([][]string, error) {
	tickers, err := polo.poloniex.GetTickers()
	if err != nil {
		return nil, err
	}
	polo.tickers = tickers
	filteredTickers := polo.filter()
	prices := [][]string{}
	for k, v := range filteredTickers {
		item := []string{k, v.Last.String()}
		prices = append(prices, item)
	}
	return prices, nil
}

// filter filter results by symbols provided in constructor
func (polo *Polo) filter() map[string]poloniex.Ticker {
	res := map[string]poloniex.Ticker{}
	polo.AvailableSymbols = []string{}
	for k, v := range polo.tickers {
		polo.AvailableSymbols = append(polo.AvailableSymbols, k)
		if StringInArray(k, polo.Symbols) > -1 {
			res[k] = v
		}
	}
	return res
}

// StringInArray check if string is in string array, -1 is not exists
func StringInArray(key string, array []string) int {
	for index, value := range array {
		if value == key {
			return index
		}
	}
	return -1
}
