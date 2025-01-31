package main

import (
	"net/http"
)

// home handler for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "home.tmpl")
}

// about handler for the about page
func (app *application) about(w http.ResponseWriter, r *http.Request) {

	app.render(w, http.StatusOK, "about.tmpl")
}

// officers handler for the officers page
func (app *application) officers(w http.ResponseWriter, r *http.Request) {

	app.render(w, http.StatusOK, "officers.tmpl")
}

// events handler for the events page
func (app *application) events(w http.ResponseWriter, r *http.Request) {

	app.render(w, http.StatusOK, "events.tmpl")
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {

	app.render(w, http.StatusOK, "login.tmpl")
}

func (app *application) healthz(w http.ResponseWriter, r *http.Request) {
	// return status ok
	w.WriteHeader(http.StatusOK)
}
