package main

import (
	"fmt"
	"github.com/almira-galeeva/testGrpc/internal/app/api/note_v1"
	"log"
	"net"

	desc "github.com/almira-galeeva/testGrpc/pkg/note_v1"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewNote())

	fmt.Println("Server is running on port:", port)
	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}