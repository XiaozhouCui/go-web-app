package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config, can be accessed by any parts of the app
type AppConfig struct {
	UseCache      bool                          // app preference, turn cache on/off
	TemplateCache map[string]*template.Template // cached templates
	InfoLog       *log.Logger                   // logger
}
