package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/ikotan/webdbpress-grpc/uploader/proto"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[file] ")
}

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s\n", err)
	}
	defer conn.Close()

	name := os.Args[1]

	fs, err := os.Open(name)
	if err != nil {
		log.Fatalf("could not open file %s\n", err)
	}
	defer fs.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(), 3*time.Second)
	defer cancel()

	c := pb.NewFileServiceClient(conn)
	stream, err := c.Upload(ctx)
	if err != nil {
		log.Fatalf("could not upload file: %s\n", err)
	}
	buf := make([]byte, 1000*1024)
	for {
		n, err := fs.Read(buf)
		println("###########")
		println(n)
		println("###########")
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %s\n", err)
		}
		stream.Send(&pb.FileRequest{
			Name: name,
			Data: buf[:n],
		})
	}
	res, err := stream.CloseAndRecv()
	log.Printf("done %d bytes\n", res.GetSize())
}
