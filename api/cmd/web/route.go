package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (app *app) routes() http.Handler {

	mux := http.NewServeMux()
	mux.Handle("GET /metrics", promhttp.Handler())
	mux.HandleFunc("GET /healthcheck", app.healthCheck)
	mux.HandleFunc("GET /{short_code}", app.GetUrl)
	mux.HandleFunc("POST /shorten", app.HandleShorten)

	return app.enableCORS(app.panicRecovery(app.rateLimit(mux)))
}
