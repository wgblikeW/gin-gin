package log

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
)

type Logger interface {
	Warningf(string, []interface{}, ...attribute.KeyValue)
	Errorf(string, []interface{}, ...attribute.KeyValue)
	Fatalf(string, []interface{}, ...attribute.KeyValue)
	Infof(string, []interface{}, ...attribute.KeyValue)
	Panicf(string, []interface{}, ...attribute.KeyValue)
	WithFields(logrus.Fields) *logrus.Entry
}
