package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/p1nant0m/gin-gin/pkg/core"
	"github.com/sirupsen/logrus"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type GenericAPIServer struct {
	TracerProvider *TraceProviderSettings

	middlewares []string
	*gin.Engine
	healthz         bool
	enableMetrics   bool
	enableProfiling bool
	enableTracing   bool
}

func initGenericAPIServer(s *GenericAPIServer) {
	s.Setup()
	s.InstallMiddlewares()
	s.InstallDefaultRoutes()
}

func (s *GenericAPIServer) Setup() {

}

func (s *GenericAPIServer) InstallMiddlewares() {
	// Default GIN middlewares
	s.Use(gin.Logger())
	s.Use(gin.Recovery())
}

func (s *GenericAPIServer) InstallDefaultRoutes() {
	// install healthz probe if it is enable
	if s.healthz {
		s.GET("/healthz", func(c *gin.Context) {
			core.WriteResponse(c, nil, gin.H{"status": "ok"})
		})
	}

	// install metrics handler if it is enable
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}

	// install pprof handler if it is enable
	if s.enableProfiling {
		pprof.Register(s.Engine)
	}

	if s.enableTracing {
		registedTracingOrExit(s.Engine, s.TracerProvider)
	}
}

func registedTracingOrExit(e *gin.Engine, cfg *TraceProviderSettings) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.ToEndPoint)))
	if err != nil {
		logrus.Panicf("error occurs when creating jaeger exporter with endpoint %v err: %v", cfg.ToEndPoint, err)
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.ServiceName),
			attribute.String("gin.realese", gin.Mode()),
		)),
	)

	e.Use(otelgin.Middleware(cfg.ServiceName, otelgin.WithTracerProvider(tp)))
}
