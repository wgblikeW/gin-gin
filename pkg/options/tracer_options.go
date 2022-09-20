package options

import "github.com/p1nant0m/gin-gin/pkg/server"

const (
	SERVICENAME = "authz-server" // DO NOT BIND FLAG
)

type TracerOptions struct {
	ToEndPoint  string `json:"endpoint" mapstructure:"endpoint"`
	ServiceName string `json:"serviceName" mapstructure:"serviceName"`
}

func NewTraceOptions() *TracerOptions {
	return &TracerOptions{
		ToEndPoint:  "http://localhost:14268/api/traces",
		ServiceName: SERVICENAME,
	}
}

func (t *TracerOptions) ApplyTo(cfg *server.Config) error {
	cfg.TraceProvider = &server.TraceProviderSettings{
		ToEndPoint:  t.ToEndPoint,
		ServiceName: t.ServiceName,
	}

	return nil
}
