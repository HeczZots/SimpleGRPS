package main

import (
	"context"
	"gRPC/internal/api/caches"
	"gRPC/internal/api/handlers"
	pb "gRPC/internal/api/proto"
	"log"
	"sync"
)

func authenticate(client pb.DataServiceClient, l, p string) (err error) {
	_, err = client.Authenticate(context.Background(), &pb.AuthRequest{
		Login:    l,
		Password: p,
	})
	return err
}

func startStream(client pb.DataServiceClient, ts int32, l string) (err error) {
	stream, err = client.StartServer(context.Background(), &pb.DataRequest{
		IntervalMs: ts,
		Login:      l,
	})
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
