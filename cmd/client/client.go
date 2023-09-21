/*
client
*/
package main

import (
	"context"
	"fmt"
	"log"
	"sync"

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

func main() {
	p = GetParams()
	//без grpc.WithInsecure()
	conn, err := grpc.Dial(p.Host+p.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось установить соединение: %v", err)
	}
	defer conn.Close()

	client = pb.NewDataServiceClient(conn)

	authRequest := &pb.AuthRequest{
		Login:    p.Login,
		Password: p.Password,
	}
	_, err = client.Authenticate(context.Background(), authRequest)
	if err != nil {
		log.Fatalf("Ошибка при аутентификации: %v", err)
	}

	dataRequest := &pb.DataRequest{
		IntervalMs: int32(p.TS),
	}
	stream, err := client.StartServer(context.Background(), dataRequest)
	if err != nil {
		log.Fatalf("Ошибка при начале передачи данных: %v", err)
	}

	select {
	case <-stream.Context().Done():
		stream.CloseSend()
	default:
		receiver = handlers.NewReceiver()
		buffer = caches.NewBuffer(p.BufferCapacity)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go receiver.GetData(buffer, stream, wg)
		wg.Wait()
	}

	for _, v := range buffer.Arr {
		fmt.Printf("Value: %v \tTime: %v\n", v.Value, v.TS.Format("15:04:05.000"))
	}
}
