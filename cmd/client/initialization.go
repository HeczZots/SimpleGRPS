package main

import (
	"context"
	"flag"
	"gRPC/internal/api/caches"
	"gRPC/internal/api/handlers"
	pb "gRPC/internal/api/proto"
	"log"
	"sync"
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

func getParams() *Flags {
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

func authenticate(client pb.DataServiceClient, l, p string) (err error) {
	authRequest := &pb.AuthRequest{
		Login:    l,
		Password: p,
	}
	_, err = client.Authenticate(context.Background(), authRequest)
	return err
}

func startStream(client pb.DataServiceClient, ts int32) (err error) {
	dataRequest := &pb.DataRequest{
		IntervalMs: int32(p.TS),
	}
	stream, err = client.StartServer(context.Background(), dataRequest)
	if err != nil {
		return err
	}
	log.Println("Stream started")
	select {
	case <-stream.Context().Done():
		stream.CloseSend()
	default:
		receiver = handlers.NewReceiver()
		buffer = caches.NewBuffer(p.BufferCapacity)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go receiver.GetData(p.TTL, buffer, stream, wg)
		wg.Wait()
	}
	return nil
}

// func Parse(p Flags){
// 	log.Printf(format, p.)
// }
