package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config, can be accessed by any parts of the app
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
