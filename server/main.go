package main

import (
  "context"
  "flag"
  "fmt"
  p "github.com/hainguyen27798/go-grpc-demo/proto"
  "google.golang.org/grpc"
  "log"
  "net"
)

var port = flag.Int("port", 9000, "The port to listen on")

type server struct {
	p.UnimplementedMailServiceServer
}

func (s *server) Send(ctx context.Context, in *p.SendMailRequest) (*p.SendMailResponse, error) {
	log.Printf("Received request to send email to: %v", in.To)
	return &p.SendMailResponse{
		StatusCode: 200,
		Message:    "Email sent successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
    grpc.MaxRecvMsgSize(1024*1024*20), // 20MB receive limit
    grpc.MaxSendMsgSize(1024*1024*20), // 20MB send limit
  )
	p.RegisterMailServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
