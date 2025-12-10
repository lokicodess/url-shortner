package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lokicodess/url-shortner/internal/data"
	"github.com/lokicodess/url-shortner/internal/model"
)

func (app app) healthCheck(w http.ResponseWriter, r *http.Request) {
	env := data.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"enviornment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app app) GetUrl(w http.ResponseWriter, r *http.Request) {
	short_code := r.PathValue("short_code")

	actual_url, err := app.urlModel.Get(short_code)

	if err != nil {
		if errors.Is(err, model.ErrRowNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}
	http.Redirect(w, r, actual_url, http.StatusPermanentRedirect)
}

func (app app) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	actualURL := r.FormValue("url")

	// validation need to be done

	shortCode := app.generateShortCode(actualURL)

	_, err := app.urlModel.Get(shortCode)

	// kinda like local caching
	if err == nil {
		shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)
		obj := data.URL{
			ShortUrl:  shortURL,
			ShortCode: shortCode,
		}
		app.writeJSON(w, 200, data.Envelope{"url": obj}, nil)
	}

	if err != model.ErrRowNotFound {
		app.notFoundResponse(w, r)
	}

	if err := app.urlModel.Post(shortCode, actualURL, 7); err != nil {
		app.serverErrorResponse(w, r, err)
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

	obj := data.URL{
		ShortUrl:  shortURL,
		ShortCode: shortCode,
	}

	err = app.writeJSON(w, 200, data.Envelope{"url": obj}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
