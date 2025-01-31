package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	// create a new logger for error messages
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// create a new logger for info messages
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// create a new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorlog.Fatal(err)
	}

	// create a new application struct
	app := NewApplication(errorlog, infolog, templateCache)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}

	infolog.Printf("Starting server on %s", *addr)
	// err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	err = srv.ListenAndServe()
	errorlog.Fatal(err)
}
