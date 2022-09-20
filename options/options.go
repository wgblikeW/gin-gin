package options

import genericOptions "github.com/p1nant0m/gin-gin/pkg/options"

type Options struct {
	TraceExporterOptions    *genericOptions.TracerOptions    `json:"exporter" mapstructure:"exporter"`
	GenericServerRunoptions *genericOptions.ServerRunOptions `json:"server"   mapstructure:"server"`
}
