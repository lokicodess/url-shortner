package main

import "net/http"

func (app *app) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{short_code}", app.GetUrl)
	mux.HandleFunc("POST /shorten", app.HandleShorten)

	return mux
}
