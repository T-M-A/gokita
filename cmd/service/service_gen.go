package service

import (
	endpoint "github.com/SeamPay/gokita/pkg/endpoint"
	service "github.com/SeamPay/gokita/pkg/service"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	grpc "github.com/go-kit/kit/transport/grpc"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initGRPCHandler(endpoints, g)
	return g
}

func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer, zipkinTracer *stdzipkin.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"ListCustomer": {grpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "ListCustomer", logger))},
	}

	if zipkinTracer != nil {
		for key, value := range options {
			value = append(value, zipkin.GRPCServerTrace(zipkinTracer, zipkin.Name(key)))
		}
	}

	return options
}

func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["ListCustomer"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ListCustomer")), endpoint.InstrumentingMiddleware(duration.With("method", "ListCustomer"))}
}

func addDefaultServiceMiddleware(logger log.Logger, counter metrics.Counter, latency metrics.Histogram, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger), service.InstrumentingMiddleware(counter, latency))
}

func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"ListCustomer"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
