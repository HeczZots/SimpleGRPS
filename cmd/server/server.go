/*
server
*/
package main

import (
	"context"
	"flag"
	"fmt"
	pb "gRPC/internal/api/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Реализация gRPC-сервиса.
type dataServiceServer struct {
	pb.UnimplementedDataServiceServer
}

// Метод для аутентификации клиента.
func (s *dataServiceServer) Authenticate(ctx context.Context, in *pb.AuthRequest) (*emptypb.Empty, error) {
	fmt.Printf("Получен запрос на аутентификацию. Логин: %s, Пароль: %s\n", in.GetLogin(), in.GetPassword())
	return &emptypb.Empty{}, nil
}

// Метод для начала передачи данных от сервера клиенту.
func (s *dataServiceServer) StartServer(in *pb.DataRequest, stream pb.DataService_StartServerServer) error {
	interval := time.Duration(in.GetIntervalMs()) * time.Millisecond
	for i := int64(1); ; i++ {
		select {
		case <-stream.Context().Done():
			log.Printf("Клиент закрыл стрим передачи данных.")
			return nil
		case <-time.After(interval):
			// Отправляем данные клиенту.
			response := &pb.DataResponse{Value: i}
			if err := stream.Send(response); err != nil {
				return err
			}
		}
	}
}

// Метод для остановки передачи данных от сервера клиенту.
func (s *dataServiceServer) StopData(ctx context.Context, in *pb.StopRequest) (*emptypb.Empty, error) {
	// Реализация остановки передачи данных, если необходимо.
	return &emptypb.Empty{}, nil
}

func main() {
	listenPort := flag.String("p", ":8080", "enter server port")
	flag.Parse()
	lis, err := net.Listen("tcp", *listenPort)
	if err != nil {
	}
	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, &dataServiceServer{})

	log.Printf("Сервер запущен на %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
