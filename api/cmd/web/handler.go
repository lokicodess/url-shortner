package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lokicodess/url-shortner/internal/model"
)

func (app app) GetUrl(w http.ResponseWriter, r *http.Request) {
	short_code := r.PathValue("short_code")
	// validation Required
	// ------------------
	actual_url, err := app.urlModel.Get(short_code)
	if err != nil {
		if errors.Is(err, model.ErrRowNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, actual_url, http.StatusPermanentRedirect)
}

func (app app) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	actualURL := r.FormValue("url")
	if actualURL == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	shortCode := app.generateShortCode(actualURL)

	_, err := app.urlModel.Get(shortCode)

	if err == nil {
		shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"short_url": "%s", "short_code": "%s"}`, shortURL, shortCode)
		return
	}

	if err != model.ErrRowNotFound {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := app.urlModel.Post(shortCode, actualURL, 7); err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"short_url": "%s", "short_code": "%s"}`, shortURL, shortCode)
}
