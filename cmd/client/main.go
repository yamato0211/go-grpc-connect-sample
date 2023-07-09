package main

import (
	"context"
	"go-grpc-connect-sample/proto/proto"
	"go-grpc-connect-sample/proto/proto/protoconnect"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
)

const origin = "http://localhost:8080"

func main() {
	ctx := context.Background()

	client := protoconnect.NewGreetServiceClient(http.DefaultClient, origin, connect.WithGRPC())
	res, err := client.Greet(ctx, connect.NewRequest(&proto.GreetRequest{
		Name: "ottotto",
	}))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
