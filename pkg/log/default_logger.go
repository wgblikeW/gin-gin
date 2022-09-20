package log

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
)

type defaultLogger struct {
	logger *logrus.Logger
}

func (s *defaultLogger) Warningf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logger.Warningf(format, args...)
}

func (s *defaultLogger) Errorf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logger.Errorf(format, args...)
}

func (s *defaultLogger) Fatalf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logger.Fatalf(format, args...)
}

func (s *defaultLogger) Infof(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logger.Infof(format, args...)
}

func (s *defaultLogger) Panicf(format string, args []interface{}, attributes ...attribute.KeyValue) {
	s.logger.Panicf(format, args...)
}

func (s *defaultLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return s.logger.WithFields(fields)
}
