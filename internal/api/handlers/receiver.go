package handlers

import (
	"fmt"
	"gRPC/internal/api/caches"
	pb "gRPC/internal/api/proto"
	"log"
	"sync"
	"time"
)

type Controler struct {
}

func NewReceiver() *Controler {
	return &Controler{}
}

// Принимаем и обрабатываем данные от сервера.
func (c *Controler) GetData(buffer *caches.Buffer, stream pb.DataService_StartServerClient, wg *sync.WaitGroup) {
	for i := buffer.Capacity; i > 0; i-- {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Ошибка при получении данных: %v", err)
			break
		}
		fmt.Printf("Получено: %v\n", response.Value)
		buffer.Insert(response.Value, time.Now())
		// В этом месте можно добавить логику обработки полученных данных.
	}
	wg.Done()
}
