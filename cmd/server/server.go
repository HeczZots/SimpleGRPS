/*
server
*/
package main

import (
	"flag"
	"gRPC/internal/api/handlers"
	pb "gRPC/internal/api/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const defaultPort = ":8080"

var lis net.Listener
var s *grpc.Server
var listenPort *string

func main() {
	listenPort := flag.String("p", defaultPort, "enter server port")
	flag.Parse()

	lis, err := net.Listen("tcp", *listenPort)
	if err != nil {
	}

	s = grpc.NewServer()
	pb.RegisterDataServiceServer(s, handlers.NewService())

	log.Printf("Server started at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error launch server: %v", err)
	}
}
