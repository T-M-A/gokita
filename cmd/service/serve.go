package service

import (
	"os"

	"github.com/SeamPay/gokita/pkg/endpoint"
	"github.com/SeamPay/gokita/pkg/service"

	"github.com/go-kit/kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	// Create a single logger, which we'll use and give to other components.
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	logger.Log("environment", environment)

	// Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	if zipkinAddr != "" {
		logger.Log("tracer", "Zipkin", "type", "OpenTracing", "URL", zipkinAddr)

		var (
			err         error
			hostPort    = "localhost:80"
			serviceName = "lorem"
			reporter    = zipkinhttp.NewReporter(zipkinAddr)
		)
		defer reporter.Close()
		zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
		zipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP))
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}

		tracer = zipkinot.Wrap(zipkinTracer)
	} else {
		logger.Log("tracer", "none")
		tracer = opentracinggo.GlobalTracer()
	}

	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

	return nil
}
