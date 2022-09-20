package tracer

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	Tracer           trace.Tracer
	ExporterSettings *ExporterSettings
}

func initTracer(t *Tracer) {
	exp, _ := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(t.ExporterSettings.ToEndPoint)))

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("auth-server"),
			attribute.String("environment", "production"),
			attribute.Int64("ID", 1),
		)),
	)

	t.Tracer = tp.Tracer(t.ExporterSettings.TracerName)
}
