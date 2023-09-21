package main

import (
	"flag"
	"time"
)

type Flags struct {
	Host           string
	Port           string
	TS             int64
	BufferCapacity int
	TTL            time.Duration
	Login          string
	Password       string
}

func GetParams() *Flags {
	host := flag.String("h", "localhost", "enter host url")
	port := flag.String("p", ":8080", "enter server port")
	login := flag.String("login", "admin", "enter your login")
	pass := flag.String("pass", "admin", "enter your password")
	ts := flag.Int64("t", 100, "enter frequency in ms")
	b := flag.Int("b", 100, "enter buffer capacity")
	ttl := flag.Duration("ttl", time.Second*10, "enter time to live conn")
	flag.Parse()
	return &Flags{
		Host:           *host,
		Port:           *port,
		TS:             *ts,
		TTL:            *ttl,
		BufferCapacity: *b,
		Login:          *login,
		Password:       *pass,
	}
}

// func Parse(p Flags){
// 	log.Printf(format, p.)
// }
