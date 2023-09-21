/*
client
*/
package main

import (
	"context"
	"fmt"
	"log"

	pb "gRPC/internal/api/proto"

	"google.golang.org/grpc"
)

func main() {
	p := GetParams()
	conn, err := grpc.Dial(p.Host+p.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось установить соединение: %v", err)
	}
	defer conn.Close()
	// go func(){
	// add interupt
	// }
	client := pb.NewDataServiceClient(conn)

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
	defer stream.CloseSend()

	// Принимаем и обрабатываем данные от сервера.
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Ошибка при получении данных: %v", err)
			break
		}
		fmt.Printf("Получено: %v\n", response.Value)

		// В этом месте можно добавить логику обработки полученных данных.
	}
}
