// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/greet.proto

package protoconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	proto "go-grpc-connect-sample/proto/proto"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "api.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServiceGreetProcedure is the fully-qualified name of the GreetService's Greet RPC.
	GreetServiceGreetProcedure = "/api.GreetService/Greet"
)

// GreetServiceClient is a client for the api.GreetService service.
type GreetServiceClient interface {
	Greet(context.Context, *connect_go.Request[proto.GreetRequest]) (*connect_go.Response[proto.GreetResponse], error)
}

// NewGreetServiceClient constructs a client for the api.GreetService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreetServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greetServiceClient{
		greet: connect_go.NewClient[proto.GreetRequest, proto.GreetResponse](
			httpClient,
			baseURL+GreetServiceGreetProcedure,
			opts...,
		),
	}
}

// greetServiceClient implements GreetServiceClient.
type greetServiceClient struct {
	greet *connect_go.Client[proto.GreetRequest, proto.GreetResponse]
}

// Greet calls api.GreetService.Greet.
func (c *greetServiceClient) Greet(ctx context.Context, req *connect_go.Request[proto.GreetRequest]) (*connect_go.Response[proto.GreetResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// GreetServiceHandler is an implementation of the api.GreetService service.
type GreetServiceHandler interface {
	Greet(context.Context, *connect_go.Request[proto.GreetRequest]) (*connect_go.Response[proto.GreetResponse], error)
}

// NewGreetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreetServiceHandler(svc GreetServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	greetServiceGreetHandler := connect_go.NewUnaryHandler(
		GreetServiceGreetProcedure,
		svc.Greet,
		opts...,
	)
	return "/api.GreetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreetServiceGreetProcedure:
			greetServiceGreetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) Greet(context.Context, *connect_go.Request[proto.GreetRequest]) (*connect_go.Response[proto.GreetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.GreetService.Greet is not implemented"))
}
