package main

import (
	authserver "github.com/p1nant0m/gin-gin"
	"github.com/p1nant0m/gin-gin/config"
	severoption "github.com/p1nant0m/gin-gin/options"
	"github.com/p1nant0m/gin-gin/pkg/options"
	"github.com/sirupsen/logrus"
)

func main() {
	// Without app package to initialize a CLI application with given flag or parse options from config files
	// we just manually configure the settings.
	opts := severoption.Options{
		TraceExporterOptions:    options.NewTraceOptions(),
		GenericServerRunoptions: options.NewServerRunOptions(),
	}

	cfg, _ := config.CreateConfigFromOptions(&opts)

	server, err := authserver.CreateAuthzServer(cfg)
	if err != nil {
		logrus.Fatalf("error occurs when creating authzServer err:%+#v", err)
	}

	server.PrepareRun().Run()
}
