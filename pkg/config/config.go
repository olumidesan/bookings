package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template // Holds template cache
	UseCache      bool
	InfoLog       *log.Logger
}
