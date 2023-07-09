package greet

import (
	"context"
	"fmt"
	"go-grpc-connect-sample/pkg/log"
	"go-grpc-connect-sample/proto/proto"
	"go-grpc-connect-sample/proto/proto/protoconnect"

	"github.com/bufbuild/connect-go"
)

type server struct {
	logger log.Handler
}

func New(logger log.Handler) protoconnect.GreetServiceHandler {
	return &server{
		logger: logger,
	}
}

func (s *server) Greet(ctx context.Context, req *connect.Request[proto.GreetRequest]) (*connect.Response[proto.GreetResponse], error) {
	s.logger.InfoCtx(ctx, "greet", "name", req.Msg.GetName())
	return connect.NewResponse(&proto.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.GetName()),
	}), nil
}
