package main

import (
	"html/template"
	"log"
)

// use a struct to hold application-wide dependencies
// this way our handlers (now methods against the application struct) can access things like loggers
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

// NewApplication creates a new application struct
func NewApplication(errorLog, infoLog *log.Logger, templateCache map[string]*template.Template) *application {
	return &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}
}
