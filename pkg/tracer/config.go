package tracer

// Config is a structure used to configure a Tracer.
type Config struct {
	ExporterSettings *ExporterSettings
}

type ExporterSettings struct {
	ToEndPoint string
	TracerName string
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{}
}

type CompletedConfig struct {
	*Config
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) New() (*Tracer, error) {
	t := &Tracer{
		ExporterSettings: c.ExporterSettings,
	}

	initTracer(t)

	return t, nil
}
