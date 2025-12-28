package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/lokicodess/url-shortner/internal/data"
	"github.com/lokicodess/url-shortner/internal/model"
	"github.com/lokicodess/url-shortner/internal/validator"
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
	shortCode := r.PathValue("short_code")

	if !validator.Matches(shortCode, validator.ShortCodeRX) {
		app.notFoundResponse(w, r)
		return
	}

	actualURL, err := app.urlModel.Get(shortCode)
	if err != nil {
		if errors.Is(err, model.ErrRowNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if !validator.Matches(actualURL, validator.UrlRX) {
		app.serverErrorResponse(w, r, errors.New("invalid stored url"))
		return
	}

	http.Redirect(w, r, actualURL, http.StatusPermanentRedirect)
}

func (app app) HandleShorten(w http.ResponseWriter, r *http.Request) {
	// limit size to prevent abuse
	r.Body = http.MaxBytesReader(w, r.Body, 1024)

	err := r.ParseMultipartForm(1024)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	actualURL := strings.TrimSpace(r.FormValue("url"))

	// --------------------------
	// VALIDATION
	// --------------------------
	v := validator.New()
	v.CheckURL("url", actualURL)

	if !v.Valid() {
		app.failedValidationResponse(w, v.Errors)
		return
	}

	shortCode := app.generateShortCode(actualURL)

	// If short code already exists → return the existing one
	_, err = app.urlModel.Get(shortCode)
	if err == nil {
		shortURL := fmt.Sprintf("https://clck.dev/%s", shortCode)
		obj := data.URL{
			ShortUrl:  shortURL,
			ShortCode: shortCode,
		}
		app.writeJSON(w, 200, data.Envelope{"url": obj}, nil)
		return
	}

	if err != model.ErrRowNotFound {
		app.serverErrorResponse(w, r, err)
		return
	}

	if err := app.urlModel.Post(shortCode, actualURL, 7); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	shortURL := fmt.Sprintf("https://clck.dev/%s", shortCode)
	obj := data.URL{
		ShortUrl:  shortURL,
		ShortCode: shortCode,
	}

	app.writeJSON(w, 200, data.Envelope{"url": obj}, nil)
}
