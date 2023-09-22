package handlers

import (
	"context"
	pb "gRPC/internal/api/proto"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Реализация gRPC-сервиса.
type DataServiceServer struct {
	pb.UnimplementedDataServiceServer
}

func (s *DataServiceServer) Authenticate(ctx context.Context, in *pb.AuthRequest) (*emptypb.Empty, error) {
	log.Printf("Auth received. Log: %s, Pass: %s\n", in.GetLogin(), in.GetPassword())
	return &emptypb.Empty{}, nil
}

func (s *DataServiceServer) StartServer(in *pb.DataRequest, stream pb.DataService_StartServerServer) error {
	interval := time.Duration(in.GetIntervalMs()) * time.Millisecond
	for i := int64(1); ; i++ {
		select {
		case <-stream.Context().Done():
			log.Printf("Client closed connection.")
			return nil
		case <-time.After(interval):
			// Отправляем данные клиенту.
			response := &pb.DataResponse{Value: i}
			log.Println("data sended: ", response.Value)
			if err := stream.Send(response); err != nil {
				return err
			}
		}
	}
}

func (s *DataServiceServer) StopData(ctx context.Context, in *pb.StopRequest) (*emptypb.Empty, error) {
	// Реализация остановки передачи данных, если необходимо.
	return &emptypb.Empty{}, nil
}
