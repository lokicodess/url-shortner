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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//  validation Required
	// ------------------
	actual_url := r.FormValue("url")
	if actual_url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	short_code := app.generateShortCode(actual_url)

	err = app.urlModel.Post(short_code, actual_url, 7)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", short_code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"short_url": "%s", "short_code": "%s"}`, shortURL, short_code)
}
