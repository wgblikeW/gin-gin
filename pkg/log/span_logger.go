package log

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type spanLogger struct {
	logger *logrus.Logger
	span   trace.Span
}

func (s *spanLogger) Warningf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logToSpan(format, "warning", args, attributes...)
	s.logger.Warningf(format, args...)
}

func (s *spanLogger) Errorf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logToSpan(format, "error", args, attributes...)
	s.logger.Errorf(format, args...)
}

func (s *spanLogger) Fatalf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logToSpan(format, "fatal", args, attributes...)
	s.logger.Fatalf(format, args...)
}

func (s *spanLogger) Infof(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logToSpan(format, "info", args, attributes...)
	s.logger.Infof(format, args...)
}

func (s *spanLogger) Panicf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logToSpan(format, "panic", args, attributes...)
	s.logger.Panicf(format, args...)
}

func (s *spanLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return s.logger.WithFields(fields)
}

func (s *spanLogger) logToSpan(format string, logLevel string, args []interface{}, attributes ...attribute.KeyValue) {
	attri := append(attributes, attribute.String("logger.level", logLevel), attribute.String("logger.msg", fmt.Sprintf(format, args...)))
	s.span.AddEvent("logger-event",
		trace.WithAttributes(attri...))
}
