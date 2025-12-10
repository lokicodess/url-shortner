package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lokicodess/url-shortner/internal/data"
)

func (app app) generateShortCode(url string) string {
	hash := md5.Sum([]byte(url))
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, encoded)
	const codeLen = 7
	if len(cleaned) > codeLen {
		return cleaned[:codeLen]
	}
	return cleaned
}

func (app app) writeJSON(w http.ResponseWriter, status int, data data.Envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
