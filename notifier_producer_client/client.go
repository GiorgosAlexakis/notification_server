package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Looper2074/notification_server/notify"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	stream, err := client.Notify(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < 5; i++ {
		order := pb.Order{Product: fmt.Sprint(i), UserID: 12345, Operation: pb.OrderOperationType_INIT_ORDER, Timestamp: timestamppb.Now()}
		err := stream.Send(&order)
		if err != nil {
			log.Println("failed to receive from notifications stream")
			return
		}
		if err == io.EOF {
			break
		}
		log.Println(&order)
	}
	stream.CloseSend()
}
