package log

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Factory struct {
	logger *logrus.Logger
}

func NewFactory(logger *logrus.Logger) *Factory {
	return &Factory{logger}
}
func (f *Factory) For(ctx context.Context) Logger {

	if span := trace.SpanFromContext(ctx); span != nil {
		// Create a logger that can auto create related span while logging
		spanLogger := spanLogger{span: span, logger: f.logger}

		span.SetAttributes(attribute.String("trace_id", span.SpanContext().TraceID().String()),
			attribute.String("span_id", span.SpanContext().SpanID().String()))
		return &spanLogger
	}

	return f.Bg()
}

func (f *Factory) Bg() Logger {
	return &defaultLogger{f.logger}
}
