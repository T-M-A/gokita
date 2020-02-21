package grpc

import (
	endpoint "github.com/SeamPay/gokita/pkg/endpoint"
	pb "github.com/SeamPay/gokita/proto"
	"github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	listCustomer grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.GokitaServer {
	return &grpcServer{
		listCustomer: makeListCustomerHandler(endpoints, options["ListCustomer"]),
	}
}
