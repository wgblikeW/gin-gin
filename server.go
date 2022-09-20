package authserver

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/p1nant0m/gin-gin/config"
	"github.com/p1nant0m/gin-gin/pkg/server"
	genericAPIServer "github.com/p1nant0m/gin-gin/pkg/server"
	"github.com/sirupsen/logrus"
)

type authzServer struct {
	genericAPIServer *server.GenericAPIServer
}

type preparedAuthzServer struct {
	*authzServer
}

func CreateAuthzServer(cfg *config.Config) (*authzServer, error) {
	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericAPIServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	authzServer := &authzServer{
		genericAPIServer: genericAPIServer,
	}

	return authzServer, nil
}

func (s *authzServer) PrepareRun() preparedAuthzServer {
	initRouter(s.genericAPIServer.Engine)

	return preparedAuthzServer{s}
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericAPIServer.Config, lastErr error) {
	genericConfig = genericAPIServer.NewConfig()

	if lastErr = cfg.GenericServerRunoptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	if lastErr = cfg.TraceExporterOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}

func (s preparedAuthzServer) Run() error {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}

	return s.genericAPIServer.Engine.Run(":" + port)
}
