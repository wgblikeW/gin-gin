package config

import "github.com/p1nant0m/gin-gin/options"

type Config struct {
	*options.Options
}

func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
