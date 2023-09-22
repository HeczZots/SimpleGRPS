package handlers

import (
	"context"
	"gRPC/internal/api/db"
	pb "gRPC/internal/api/proto"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Реализация gRPC-сервиса.
type DataServiceServer struct {
	pb.UnimplementedDataServiceServer
	Users *db.Users
}

func NewService() *DataServiceServer {
	return &DataServiceServer{Users: db.NewDataBase()}
}
func (s *DataServiceServer) Authenticate(ctx context.Context, in *pb.AuthRequest) (*emptypb.Empty, error) {
	user := in.GetLogin()
	pass := in.GetPassword()
	ok := s.Users.AddUser(user, pass)
	if !ok {
		log.Printf("Auth already done. Login: %s", user)
		return nil, nil
	}
	log.Printf("Auth received. Log: %s, Pass: %s\n", user, pass)
	return &emptypb.Empty{}, nil
}

func (s *DataServiceServer) StartServer(in *pb.DataRequest, stream pb.DataService_StartServerServer) error {
	interval := time.Duration(in.GetIntervalMs()) * time.Millisecond
	for i := int64(1); ; i++ {
		select {
		case <-stream.Context().Done():
			log.Printf("Client %v closed connection", in.Login)
			s.Users.CloseUserConnection(in.Login)
			return nil
		case <-time.After(interval):
			// Отправляем данные клиенту.
			response := &pb.DataResponse{Value: i}
			log.Printf("data sended:%v to %v", response.Value, in.Login)
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
