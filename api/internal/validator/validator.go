package validator

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	UrlRX       = regexp.MustCompile(`^(https?://)([a-zA-Z0-9.-]+)(:[0-9]{1,5})?(/.*)?$`)
	ShortCodeRX = regexp.MustCompile(`^[A-Za-z0-9_-]{5,20}$`)
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func Matches(values string, rx *regexp.Regexp) bool {
	return rx.MatchString(values)
}

func IsPrivateHost(host string) bool {
	privatePrefixes := []string{
		"127.", "10.", "192.168.", "172.16.", "172.17.", "172.18.", "172.19.",
		"172.20.", "172.21.", "172.22.", "172.23.", "172.24.", "172.25.",
		"172.26.", "172.27.", "172.28.", "172.29.", "172.30.", "172.31.",
		"localhost",
	}

	for _, p := range privatePrefixes {
		if strings.HasPrefix(host, p) {
			return true
		}
	}
	return false
}

func (v *Validator) CheckURL(field, value string) {
	v.Check(value != "", field, "url is required")
	if value == "" {
		return
	}

	v.Check(len(value) <= 2000, field, "url too long")

	v.Check(UrlRX.MatchString(value), field, "invalid url format")

	u, err := url.ParseRequestURI(value)
	v.Check(err == nil, field, "invalid url")
	if err != nil {
		return
	}

	host := u.Hostname()
	v.Check(!IsPrivateHost(host), field, "private/internal urls not allowed")
}
