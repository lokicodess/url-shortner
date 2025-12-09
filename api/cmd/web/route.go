package main

import "net/http"

func (app *app) routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{short_code}", app.GetUrl)
	mux.HandleFunc("POST /shorten", app.HandleShorten)

	return app.enableCORS(mux)
}
