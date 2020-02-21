package service

import (
	"context"
	"fmt"
	pb "github.com/SeamPay/gokita/proto"
	log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"time"
)

type Middleware func(GokitaService) GokitaService

type loggingMiddleware struct {
	logger log.Logger
	next   GokitaService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a GokitaService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next GokitaService) GokitaService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (res *pb.ListCustomerReply, err error) {
	defer func() {
		l.logger.Log("method", "ListCustomer", "req", req, "res", res, "err", err)
	}()
	return l.next.ListCustomer(ctx, req)
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           GokitaService
}

// InstrumentingMiddleware takes a logger as a dependency
// and returns a GokitaService Middleware.
func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) Middleware {
	return func(next GokitaService) GokitaService {
		return &instrumentingMiddleware{requestCount, requestLatency, next}
	}
}

func (i instrumentingMiddleware) ListCustomer(ctx context.Context, req *pb.ListCustomerRequest) (res *pb.ListCustomerReply, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "ListCustomer", "error", fmt.Sprint(err != nil)}
		i.requestCount.With(lvs...).Add(1)
		i.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ListCustomer(ctx, req)
}
