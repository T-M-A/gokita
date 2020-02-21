package endpoint

import (
	"github.com/SeamPay/gokita/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	ListCustomerEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.GokitaService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		ListCustomerEndpoint: MakeListCustomerEndpoint(s),
	}
	for _, m := range mdw["ListCustomer"] {
		eps.ListCustomerEndpoint = m(eps.ListCustomerEndpoint)
	}
	return eps
}
