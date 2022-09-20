package log

import (
	"github.com/p1nant0m/gin-gin/pkg/log"
	"github.com/sirupsen/logrus"
)

var (
	defaultLogger log.Factory = *log.NewFactory(logrus.New())
)

func Logger() *log.Factory {
	return &defaultLogger
}

func Warp(items ...interface{}) (ret []interface{}) {
	if items == nil {
		return
	}

	ret = append(ret, items...)
	return
}
