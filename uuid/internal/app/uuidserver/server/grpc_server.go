package server

import (
	uuidv1 "github.com/douyu/jupiter-examples/uuid/gen/api/go/uuid/v1"
	"github.com/douyu/jupiter-examples/uuid/internal/app/uuidserver/controller"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
)

// var GrpcProviderSet = wire.NewSet(NewGrpcServer)

type GrpcServer struct {
	*xgrpc.Server
	controller.Options
}

func NewGrpcServer(opts controller.Options) *GrpcServer {
	server := xgrpc.StdConfig("grpc").MustBuild()
	uuidv1.RegisterUuidServiceServer(server.Server, opts.UuidGrpc)
	return &GrpcServer{
		Server: server,
	}
}
