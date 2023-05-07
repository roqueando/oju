package parser

import (
	"errors"
	"oluwoye/internal/config"
	"strings"
)

type Request struct {
	Header  Header
	Timer   string
	Message string
}

type Header struct {
	Verb, AppKey, Version string
}

func NewHeader(head string, allowed_applications []config.Application) (Header, error) {
	parts := strings.Split(head, " ")
	if len(parts) != 3 {
		return Header{}, errors.New("malformed header")
	}

	verb := parts[0]
	app_key := parts[1]
	version := parts[2]

	if !is_verb_allowed(verb) {
		return Header{}, errors.New("verb not allowed")
	}

	if !is_application_allowed(app_key, allowed_applications) {
		return Header{}, errors.New("application not allowed")
	}

	return Header{Verb: verb, AppKey: app_key, Version: version}, nil
}

func is_verb_allowed(verb string) bool {
	allowed := []string{
		"LOG",
		"QUERY",
	}

	for _, allowed_verb := range allowed {
		if allowed_verb == verb {
			return true
		}
	}
	return false
}

func is_application_allowed(app_key string, allowed_applications []config.Application) bool {
	is_allowed := false

	for _, application := range allowed_applications {
		if application.AppKey == app_key {
			is_allowed = true
		} else {
			is_allowed = false
		}
	}
	return is_allowed
}