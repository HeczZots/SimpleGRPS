package config

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

const (
	defaultPort       = ":8080"
	defaultHost       = "localhost"
	defaultLogin      = "admin"
	defaultPassword   = "admin"
	defaultTimeLaps   = 100
	defaultTimeToLive = time.Second * 1
	defaultCapacity   = 50
)

func ParseFlags() *Flags {
	host := flag.String("h", defaultHost, "enter host url")
	port := flag.String("p", defaultPort, "enter server port")
	login := flag.String("login", defaultLogin, "enter your login")
	pass := flag.String("pass", defaultPassword, "enter your password")
	ts := flag.Int64("t", defaultTimeLaps, "enter frequency in ms")
	b := flag.Int("b", defaultCapacity, "enter buffer capacity")
	ttl := flag.Duration("ttl", defaultTimeToLive, "enter time to live conn")
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
