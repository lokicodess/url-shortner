package main

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

func (app *app) generateShortCode(url string) string {
	hash := md5.Sum([]byte(url))
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	cleaned := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, encoded)
	return cleaned
}
