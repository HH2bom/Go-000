package rpc

import (
	"context"

	"google.golang.org/grpc"

	"app/internal/pkg/conf"
)

func Server(ctx context.Context, cfg conf.C) (*grpc.Server, func() error, func()) {
	server := grpc.NewServer()
	runServer := func() error {
		return server.Serve(nil)
	}

	stopServer := func() {
		server.GracefulStop()
	}

	return server, runServer, stopServer
}
