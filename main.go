package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/GrantZheng/monolith_demo/common"

	barragesvc "github.com/GrantZheng/monolith_demo/barrage/cmd/service"
	giftsvc "github.com/GrantZheng/monolith_demo/gift/cmd/service"
	liveroomsvc "github.com/GrantZheng/monolith_demo/live_room/cmd/service"

	log "github.com/go-kit/kit/log"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	zipkingoopentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkingo "github.com/openzipkin/zipkin-go"
	http "github.com/openzipkin/zipkin-go/reporter/http"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	appdash "sourcegraph.com/sourcegraph/appdash"
	opentracing "sourcegraph.com/sourcegraph/appdash/opentracing"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("barrage", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
var thriftAddr = fs.String("thrift-addr", ":8083", "Thrift listen address")
var thriftProtocol = fs.String("thrift-protocol", "binary", "binary, compact, json, simplejson")
var thriftBuffer = fs.Int("thrift-buffer", 0, "0 for unbuffered")
var thriftFramed = fs.Bool("thrift-framed", false, "true to enable framing")
var zipkinURL = fs.String("zipkin-url", "", "Enable Zipkin tracing via a collector URL e.g. http://localhost:9411/api/v1/spans")
var lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
var appdashAddr = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")

func run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	common.Logger = log.NewLogfmtLogger(os.Stderr)
	common.Logger = log.With(common.Logger, "ts", log.DefaultTimestampUTC)
	common.Logger = log.With(common.Logger, "caller", log.DefaultCaller)

	// Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	if *zipkinURL != "" {
		common.Logger.Log("tracer", "Zipkin", "URL", *zipkinURL)
		reporter := http.NewReporter(*zipkinURL)
		defer reporter.Close()
		endpoint, err := zipkingo.NewEndpoint("barrage", "localhost:80")
		if err != nil {
			common.Logger.Log("err", err)
			os.Exit(1)
		}
		localEndpoint := zipkingo.WithLocalEndpoint(endpoint)
		nativeTracer, err := zipkingo.NewTracer(reporter, localEndpoint)
		if err != nil {
			common.Logger.Log("err", err)
			os.Exit(1)
		}
		common.Tracer = zipkingoopentracing.Wrap(nativeTracer)
	} else if *lightstepToken != "" {
		common.Logger.Log("tracer", "LightStep")
		common.Tracer = lightsteptracergo.NewTracer(lightsteptracergo.Options{AccessToken: *lightstepToken})
		defer lightsteptracergo.Flush(context.Background(), common.Tracer)
	} else if *appdashAddr != "" {
		common.Logger.Log("tracer", "Appdash", "addr", *appdashAddr)
		collector := appdash.NewRemoteCollector(*appdashAddr)
		common.Tracer = opentracing.NewTracer(collector)
		defer collector.Close()
	} else {
		common.Logger.Log("tracer", "none")
		common.Tracer = opentracinggo.GlobalTracer()
	}

	g := &group.Group{}
	initHttpHandler(g)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	common.Logger.Log("exit", g.Run())
}

func initMetricsEndpoint(g *group.Group) {
	http2.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		common.Logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		common.Logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http2.Serve(debugListener, http2.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initHttpHandler(g *group.Group) http2.Handler {
	mux := http2.NewServeMux()
	barragesvc.InitHttpHandler(mux)
	giftsvc.InitHttpHandler(mux)
	liveroomsvc.InitHttpHandler(mux)

	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		common.Logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		common.Logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http2.Serve(httpListener, mux)
	}, func(error) {
		httpListener.Close()
	})
	return mux
}

func main() {
	run()
}
