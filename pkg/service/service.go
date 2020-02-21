package service

import (
	"context"
	pb "github.com/SeamPay/gokita/proto"
)

// GokitaService describes the service.
type GokitaService interface {
	// Add your methods here
	ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (res *pb.ListCustomerReply, err error)
}

type basicGokitaService struct{}

func (b *basicGokitaService) ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (res *pb.ListCustomerReply, err error) {
	// TODO implement the business logic of ListCustomer

	var customers []*pb.Customer
	res = &pb.ListCustomerReply{
		Items: customers,
		Total: 0,
		Page:  0,
	}

	return res, err
}

// NewBasicGokitaService returns a naive, stateless implementation of GokitaService.
func NewBasicGokitaService() GokitaService {
	return &basicGokitaService{}
}

// New returns a GokitaService with all of the expected middleware wired in.
func New(middleware []Middleware) GokitaService {
	var svc GokitaService = NewBasicGokitaService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
