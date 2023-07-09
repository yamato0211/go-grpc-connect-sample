package main

import (
	"context"
	"errors"
	"go-grpc-connect-sample/pkg/handler/greet"
	"go-grpc-connect-sample/pkg/log"
	"go-grpc-connect-sample/proto/proto/protoconnect"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	os.Exit(run())
}

func run() int {
	const (
		ok = iota
		ng
	)

	//DI
	logger := log.NewHandler(log.LevelInfo, log.WithJSONFormat())
	greetServer := greet.New(logger)

	mux := http.NewServeMux()
	mux.Handle(protoconnect.NewGreetServiceHandler(greetServer))
	handler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorCtx(ctx, "failed to ListenAndServe", "err", err)
		}
	}()

	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		return ng
	}

	return ok
}
