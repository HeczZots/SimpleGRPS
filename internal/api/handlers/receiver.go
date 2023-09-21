package handlers

import (
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
func (c *Controler) GetData(d time.Duration, buffer *caches.Buffer, stream pb.DataService_StartServerClient, wg *sync.WaitGroup) {
	startTime := time.Now()
	for i := buffer.Capacity; i > 0; i-- {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error in getting data: %v", err)
			break
		}
		if time.Since(startTime) > d {
			log.Printf("Duration %v has expired\n", d)
			break
		}
		log.Printf("Received: %v\n", response.Value)
		buffer.Insert(response.Value, time.Now())
	}
	wg.Done()
}
