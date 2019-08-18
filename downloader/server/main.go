package main

import (
	"log"
	"net"

	pb "github.com/ikotan/webdbpress-grpc/downloader/proto"
	"google.golang.org/grpc"
)

const port string = ":50051"

func init() {
	log.SetFlags(0)
	log.SetPrefix("[file] ")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	srv := grpc.NewServer()
	pb.RegisterFileServiceServer(srv, &fileService{})
	log.Printf("start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
