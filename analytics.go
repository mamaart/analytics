package analytics

import (
	"log"
	"net/http"
)

type Logger interface {
	Log(r *http.Request, eventName string)
}

type logger struct{}

func New() Logger { return &logger{} }

func (l *logger) Log(r *http.Request, eventName string) {
	log.Println(eventName)
}
