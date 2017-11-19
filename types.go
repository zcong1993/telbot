package main

import (
	"github.com/shopspring/decimal"
)

// Subscriber can add a price notification listener
type Subscriber struct {
	// Symbol is pair
	Symbol string
	// Price is notification edge
	Price decimal.Decimal
	// Operation is ">" or "<"
	Operation string
	// Times is notification times
	Times int
}

// SymbolSubscribers is Subscriber array with helper functions
type SymbolSubscribers []Subscriber

func (ss SymbolSubscribers) findIndex(s Subscriber) int {
	for index, item := range ss {
		if item == s {
			return index
		}
	}
	return -1
}

// UpdateOrCreate append an item if not exists, otherwise update it
func (ss SymbolSubscribers) UpdateOrCreate(pre, new Subscriber) SymbolSubscribers {
	index := ss.findIndex(pre)
	if index == -1 {
		ss = append(ss, new)
	} else {
		ss[index] = new
	}
	return ss
}

// Delete will delete an array item if exists
func (ss SymbolSubscribers) Delete (s Subscriber) SymbolSubscribers {
	index := ss.findIndex(s)
	if index > -1 {
		ss = append(ss[:index], ss[index+1:]...)
	}
	return ss
}
