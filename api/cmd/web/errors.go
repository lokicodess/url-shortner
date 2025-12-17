package main

import (
	"fmt"
	"net/http"

	"github.com/lokicodess/url-shortner/internal/data"
)

func (app app) failedValidationResponse(w http.ResponseWriter, errors map[string]string) {
	envelope := data.Envelope{
		"error":   "validation_failed",
		"message": "Input validation failed",
		"fields":  errors,
	}
	app.writeJSON(w, 422, envelope, nil)
}

func (app app) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, r, http.StatusTooManyRequests, message)
}

func (app app) logError(r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app app) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := data.Envelope{"error": message}

	if err := app.writeJSON(w, status, env, nil); err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app app) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app app) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app app) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app app) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
