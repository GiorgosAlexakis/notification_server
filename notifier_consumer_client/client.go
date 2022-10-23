package main

import (
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/Looper2074/notification_server/notify"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	//subch = make(chan pb.OrderOperationType, 100)
	conn, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Println("client could not connect to server")
	}
	client := pb.NewNotifierClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rand.Seed(time.Now().UTC().UnixNano())
	stream, err := client.Subscribe(ctx, &pb.Subscriber{SubscriberID: 12345})
	if err != nil {
		log.Println(err)
		return
	}
	for {
		notification, err := stream.Recv()
		if err != nil {
			log.Println("failed to receive from notifications stream")
			return
		}
		if err == io.EOF {
			break
		}
		log.Println(notification)
	}
	stream.CloseSend()
}
