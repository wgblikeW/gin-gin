package server

type Config struct {
	Mode        string
	Middlewares []string

	Healthz bool

	EnableProfiling bool
	EnableMetrics   bool
	EnableTracing   bool
}
