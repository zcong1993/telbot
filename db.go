package main

import (
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

var client *ledis.Ledis

func init() {
	cfg := lediscfg.NewConfigDefault()
	l, err := ledis.Open(cfg)
	if err != nil {
		panic(err)
	}
	client = l
}

// NewDB create a db instance with index provided
func NewDB(index int)(db *ledis.DB, err error) {
	db, err = client.Select(index)
	return
}
