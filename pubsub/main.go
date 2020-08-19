package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	pb "github.com/micro/services/pubsub/proto"
)

// Pub will publish messages every second
func Pub() {
	ev := service.NewEvent("messages")

	for {
		ev.Publish(context.TODO(), &pb.Message{
			Id: "1",
			Body: []byte(`hello`),
		})

		time.Sleep(time.Second)
	}
}

// Sub processes messages
func Sub(ctx context.Context, msg *pb.Message) error {
	fmt.Println("Received a message")
	return nil
}

func main() {
	service := service.New(
		service.Name("pubsub"),
	)

	// subscribe to the "messages" topic
	service.Subscribe("messages", Sub)

	// publish messages
	go Pub()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}