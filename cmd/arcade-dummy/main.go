package main

import (
	"os"
	"log"
	"fmt"
	"google.golang.org/grpc"
	"github.com/chakrit/arcade"
	"net"
	"github.com/chakrit/arcade/interceptors"
)

const ListenAddr = "0.0.0.0:8884"

func main() {
	defer log.Println("stopped")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s (node-identifier)\n", os.Args[0])
		return
	}

	node := &DummyNode{identifier: os.Args[1]}
	log.Println("identifier:", node.identifier)

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptors.LogServerCalls))
	arcade.RegisterNodeServiceServer(server, node)

	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("listen on", ListenAddr)
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
