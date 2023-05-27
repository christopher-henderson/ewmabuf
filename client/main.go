package main

import (
	"context"
	"fmt"
	"github.com/whatever/repro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"sync"
)

//go:generate protoc -I ../proto --go-grpc_out=./repro --go_out=./repro  ../proto/repro.proto

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	target := "localhost:8080"
	if len(os.Args) > 1 {
		target = os.Args[1]
	}
	fmt.Printf("Target: %s\n", target)
	message := &repro.Bytes{Message: make([]byte, 1024*1024*3)}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			do(target, message)
			fmt.Println(j)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}

func do(target string, message *repro.Bytes) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := repro.NewReproClient(conn)
	stream, err := client.LargeClientStream(context.Background())
	if err != nil {
		panic(err)
	}
	defer stream.CloseSend()
	for i := 0; i < 1024/3; i++ {
		err = stream.Send(message)
		if err != nil {
			panic(err)
		}
	}
}
