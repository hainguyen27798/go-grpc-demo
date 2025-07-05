package main

import (
	"context"
	"github.com/hainguyen27798/go-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"time"
)

func main() {
	address := "localhost:9000"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to server: " + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	client := proto.NewMailServiceClient(conn)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		text := string(make([]byte, 10*1024*1024))
		r, err := client.Send(ctx, &proto.SendMailRequest{
			To:           []string{"client"},
			Subject:      "subject",
			Text:         &text,
			TemplateName: nil,
			TemplateData: nil,
		})
		if err != nil {
			panic("Failed to send mail: " + err.Error())
		}
		log.Printf("Received response from sent email: %v", protojson.Format(r))
		cancel()
	}
}
