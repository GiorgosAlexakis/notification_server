package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/Looper2074/notification_server/notify"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const successfulSubscriptionMessage = "you have subscribed to the server"
const successfulNotifyMessage = "server notified"

type orderNotifierServer struct {
	pb.UnimplementedNotifierServer
	mu            sync.Mutex
	subscriptions map[uint32]pb.Notifier_SubscribeServer
}

func newServer() *orderNotifierServer {
	return &orderNotifierServer{subscriptions: make(map[uint32]pb.Notifier_SubscribeServer)}
}

func (s *orderNotifierServer) Notify(stream pb.Notifier_NotifyServer) error {
	for {
		o, err := stream.Recv()
		if err != nil {
			log.Println(err)
			return err
		}
		var operationMessage string
		if o.GetOperation() == pb.OrderOperationType_INIT_ORDER {
			operationMessage = "initialized"
		} else {
			operationMessage = "cancelled"
		}
		message := fmt.Sprintf("Order of product id: %s, was succesfully %s", o.Product, operationMessage)
		userID := o.UserID
		notification := &pb.OrderNotification{Message: message, Timestamp: o.Timestamp}
		err = s.SendNotification(userID, notification)
		if err != nil {
			log.Println(err)
		}
		return nil
	}

}

func (s *orderNotifierServer) Subscribe(sub *pb.Subscriber, stream pb.Notifier_SubscribeServer) error {
	log.Printf("Received subscription request from user: %d\n", sub.GetSubscriberID())
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.subscriptions[sub.GetSubscriberID()]
	if !ok {
		s.subscriptions[sub.GetSubscriberID()] = stream
	} else {
		log.Printf("Found existing subscription from user: %d\n", sub.GetSubscriberID())
	}
	return nil
}

func (s *orderNotifierServer) SendNotification(subscriber uint32, notification *pb.OrderNotification) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	stream, ok := s.subscriptions[subscriber]
	if !ok {
		log.Printf("No existing subscription from user: %d, skipping notification\n", subscriber)
	} else {
		if err := stream.Send(notification); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterNotifierServer(s, myServer)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
