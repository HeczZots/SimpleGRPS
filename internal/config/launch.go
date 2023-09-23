package config

import (
	"flag"
	"time"
)

type Flags struct {
	TS             int64
	BufferCapacity int
	TTL            time.Duration
	Login          string
	Password       string
}

const (
	defaultLogin      = "admin"
	defaultPassword   = "admin"
	defaultTimeLaps   = 100
	defaultTimeToLive = time.Second * 1
	defaultCapacity   = 50
)

func ParseFlags(dHost, dPort string) (*Flags, string) {
	host := flag.String("host", dHost, "enter host")
	port := flag.String("port", dPort, "enter port in format \":5555\"")
	login := flag.String("login", defaultLogin, "enter your login")
	pass := flag.String("pass", defaultPassword, "enter your password")
	ts := flag.Int64("t", defaultTimeLaps, "enter frequency in ms")
	b := flag.Int("b", defaultCapacity, "enter buffer capacity")
	ttl := flag.Duration("ttl", defaultTimeToLive, "enter time to live conn")
	flag.Parse()
	return &Flags{
		TS:             *ts,
		TTL:            *ttl,
		BufferCapacity: *b,
		Login:          *login,
		Password:       *pass,
	}, *host + *port
}
