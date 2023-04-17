package logger

import (
	"errors"
	"strings"
)

type Log struct {
	Header  Header
	Timer   string
	Message string
}

type Header struct {
	Verb, AppKey, Version string
}

func NewHeader(head string) (Header, error) {
	parts := strings.Split(head, " ")
	if len(parts) > 3 || len(parts) < 3 {
		return Header{}, errors.New("malformed header")
	}

	verb := parts[0]
	app_key := parts[0]
	version := parts[0]

	return Header{Verb: verb, AppKey: app_key, Version: version}, nil
}
