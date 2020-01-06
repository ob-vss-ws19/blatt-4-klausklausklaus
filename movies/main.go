package main

import (
	"fmt"
	"log"
	"time"

	hello "github.com/ob-vss-ws19/blatt-4-klausklausklaus/movies/proto/hello"

	"github.com/micro/go-micro"

	"context"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))
	fmt.Println("ssssss")

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
