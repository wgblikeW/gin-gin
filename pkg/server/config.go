package server

import "github.com/gin-gonic/gin"

type Config struct {
	TraceProvider   *TraceProviderSettings
	Middlewares     []string
	Healthz         bool
	EnableMetrics   bool
	EnableProfiling bool
	EnableTracing   bool
	Mode            string
}

type TraceProviderSettings struct {
	ServiceName string
	ToEndPoint  string
}

func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.DebugMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableTracing:   true,
		EnableMetrics:   true,
	}
}

type CompletedConfig struct {
	*Config
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) New() (*GenericAPIServer, error) {
	gin.SetMode(c.Mode)

	s := &GenericAPIServer{
		TracerProvider:  c.TraceProvider,
		healthz:         c.Healthz,
		enableMetrics:   c.EnableMetrics,
		enableProfiling: c.EnableProfiling,
		enableTracing:   c.EnableTracing,
		middlewares:     c.Middlewares,
		Engine:          gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}
