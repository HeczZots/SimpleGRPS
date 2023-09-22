package handlers

import (
	"context"
	"errors"
	"gRPC/internal/api/db"
	pb "gRPC/internal/api/proto"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Реализация gRPC-сервиса.
type DataServiceServer struct {
	pb.UnimplementedDataServiceServer
	Clients *db.Clients
}

func NewService() *DataServiceServer {
	return &DataServiceServer{
		Clients: db.NewClients(),
	}
}

func (s *DataServiceServer) Authenticate(ctx context.Context, in *pb.AuthRequest) (*emptypb.Empty, error) {
	l := in.GetLogin()
	p := in.GetPassword()
	ok := s.Clients.ActiveSessions(l, p)
	if !ok {
		return nil, errors.New("auth already exist")
	}
	log.Printf("Auth received. Log: %s, Pass: %s\n", l, p)
	return &emptypb.Empty{}, nil
}

func (s *DataServiceServer) StartServer(ar *pb.AuthRequest, in *pb.DataRequest, stream pb.DataService_StartServerServer) error {
	interval := time.Duration(in.GetIntervalMs()) * time.Millisecond
	for i := int64(1); ; i++ {
		select {
		case <-stream.Context().Done():
			log.Printf("Client closed connection.")
			s.Clients.ViewActiveSessions()
			s.Clients.CloseAuth(ar.GetLogin())
			return nil
		case <-time.After(interval):
			response := &pb.DataResponse{Value: i}
			log.Println("data sended: to user: ", response.Value, ar.GetLogin())
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
