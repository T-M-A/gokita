package endpoint

import (
	"context"

	service "github.com/SeamPay/gokita/pkg/service"
	pb "github.com/SeamPay/gokita/proto"
	endpoint "github.com/go-kit/kit/endpoint"
)

// ListCustomerRequest collects the request parameters for the ListCustomer method.
type ListCustomerRequest struct {
	Req *pb.ListCustomerRequest `json:"req,omitempty"`
}

// ListCustomerResponse collects the response parameters for the ListCustomer method.
type ListCustomerResponse struct {
	Res *pb.ListCustomerReply `json:"res"`
	Err error                 `json:"err"`
}

// MakeListCustomerEndpoint returns an endpoint that invokes ListCustomer on the service.
func MakeListCustomerEndpoint(s service.GokitaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListCustomerRequest)
		res, err := s.ListCustomer(ctx, req.Req)
		return ListCustomerResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r ListCustomerResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}
