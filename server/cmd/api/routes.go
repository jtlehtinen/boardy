package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/health", app.health)
	return router
}
