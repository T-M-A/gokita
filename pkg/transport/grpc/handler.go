package grpc

import (
	"context"
	endpoint "github.com/SeamPay/gokita/pkg/endpoint"
	pb "github.com/SeamPay/gokita/proto"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeListCustomerHandler creates the handler logic
func makeListCustomerHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ListCustomerEndpoint, decodeListCustomerRequest, encodeListCustomerResponse, options...)
}

// decodeListCustomerResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain ListCustomer request.
func decodeListCustomerRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ListCustomerRequest)
	return endpoint.ListCustomerRequest{Req: req}, nil
}

// encodeListCustomerResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeListCustomerResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.ListCustomerResponse)
	return &pb.ListCustomerReply{Items: resp.Res.Items, Total: resp.Res.Total, Page: resp.Res.Page, Err: err2str(resp.Err)}, nil
}

func (g *grpcServer) ListCustomer(ctx context1.Context, req *pb.ListCustomerRequest) (*pb.ListCustomerReply, error) {
	_, rep, err := g.listCustomer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListCustomerReply), nil
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
