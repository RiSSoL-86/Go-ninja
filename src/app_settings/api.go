package app_settings

import "github.com/danielgtaylor/huma/v2"

type APISettings struct {
	Title   string
	Version string
}

func NewAPISettings() *APISettings {
	return &APISettings{
		Title:   "Go Ninja API",
		Version: "1.0.0",
	}
}

func (s *APISettings) HumaConfig() huma.Config {
	config := huma.DefaultConfig(s.Title, s.Version)
	config.DocsRenderer = huma.DocsRendererSwaggerUI
	return config
}
