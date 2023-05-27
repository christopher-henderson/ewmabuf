package main

import (
	"fmt"
	"github.com/whatever/repro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"os"
)

//go:generate protoc -I ../proto --go-grpc_out=./repro --go_out=./repro  ../proto/repro.proto

type Repro struct {
	repro.UnimplementedReproServer
}

func (r Repro) LargeClientStream(server repro.Repro_LargeClientStreamServer) error {
	var err error
	for _, err = server.Recv(); err == nil; _, err = server.Recv() {
	}
	fmt.Println(err)
	return nil
}

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	target := "localhost:8080"
	if len(os.Args) > 1 {
		target = os.Args[1]
	}
	fmt.Printf("Target: %s\n", target)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	repro.RegisterReproServer(server, &Repro{})
	err = server.Serve(listener)
	fmt.Println(err)
}
