/*
client
*/
package main

import (
	"fmt"
	"log"
	"time"

	"gRPC/internal/api/caches"
	"gRPC/internal/api/handlers"
	pb "gRPC/internal/api/proto"

	"google.golang.org/grpc"
)

const (
	defaultPort       = ":8080"
	defaultHost       = "localhost"
	defaultLogin      = "admin"
	defaultPassword   = "admin"
	defaultTimeLaps   = 100
	defaultTimeToLive = time.Second * 1
	defaultCapacity   = 50
)

var p *Flags
var conn *grpc.ClientConn
var client pb.DataServiceClient
var stream pb.DataService_StartServerClient
var receiver *handlers.Controler
var buffer *caches.Buffer

func main() {
	p = getParams()
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

func outPut(buffer *caches.Buffer) {
	for _, v := range buffer.Arr {
		fmt.Printf("Value: %v \tTime: %v\n", v.Value, v.TS.Format("15:04:05.000"))
	}
}
