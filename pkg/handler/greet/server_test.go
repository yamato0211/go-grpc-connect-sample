package greet

import (
	"context"
	"fmt"
	"go-grpc-connect-sample/pkg/log/mock_log"
	"go-grpc-connect-sample/proto/proto"
	"go-grpc-connect-sample/proto/proto/protoconnect"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	logger *mock_log.MockHandler
}

func newWithMocks(t *testing.T) (context.Context, protoconnect.GreetServiceHandler, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logger := mock_log.NewMockHandler(ctrl)
	return ctx, New(
			logger,
		), &mocks{
			logger,
		}
}

func TestServer_Greet(t *testing.T) {
	name := "ottotto"

	ctx, s, m := newWithMocks(t)
	m.logger.EXPECT().InfoCtx(ctx, "greet", "name", name)
	res, err := s.Greet(ctx, connect.NewRequest(&proto.GreetRequest{
		Name: name,
	}))
	assert.Equal(t, connect.NewResponse(&proto.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", name),
	}), res)
	assert.NoError(t, err)
}
