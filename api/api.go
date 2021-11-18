package api

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedCloudHashingOrderServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterCloudHashingOrderServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterCloudHashingOrderHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
