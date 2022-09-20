package options

import (
	"github.com/gin-gonic/gin"
	"github.com/p1nant0m/gin-gin/pkg/server"
)

type ServerRunOptions struct {
	Mode        string   `json:"mode" mapstructure:"mode"`
	Middlewares []string `josn:"middleware" mapstructure:"middlewares"`

	Healthz bool `json:"healthz" mapstructure:"healthz"`
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Mode:        gin.ReleaseMode,
		Healthz:     true,
		Middlewares: []string{},
	}
}

func (s *ServerRunOptions) ApplyTo(cfg *server.Config) error {
	cfg.Middlewares = s.Middlewares
	cfg.Healthz = s.Healthz
	cfg.Mode = s.Mode

	return nil
}
