/*
client
*/
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"gRPC/internal/api/caches"
	"gRPC/internal/api/handlers"
	pb "gRPC/internal/api/proto"

	"google.golang.org/grpc"
)

var p *Flags
var conn *grpc.ClientConn
var client pb.DataServiceClient
var stream pb.DataService_StartServerClient
var receiver *handlers.Controler
var buffer *caches.Buffer

const (
	defaultPort       = ":8080"
	defaultHost       = "localhost"
	defaultLogin      = "admin"
	defaultPassword   = "admin"
	defaultTimeLaps   = 100
	defaultTimeToLive = time.Second * 1
	defaultCapacity   = 50
)

func main() {
	p = GetParams()
	url := p.Host + p.Port
	//без grpc.WithInsecure() не работает(нужно разобраться)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error to dial connection: %v", err)
	}
	defer conn.Close()
	log.Println("Connection succesful")

	client = pb.NewDataServiceClient(conn)
	err = authenticate(client, p.Login, p.Password)
	if err != nil {
		log.Fatalf("Auth error: %v", err)
	}
	log.Println("Authentification succesful")

	err = startStream(client, int32(p.TS))
	if err != nil {
		log.Fatalf("error stream creation: %v", err)
	}

	outPut(buffer)
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
func outPut(buffer *caches.Buffer) {
	for _, v := range buffer.Arr {
		fmt.Printf("Value: %v \tTime: %v\n", v.Value, v.TS.Format("15:04:05.000"))
	}
}
