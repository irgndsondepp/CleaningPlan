package rest

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRoute(name, method, pattern string, hf http.HandlerFunc) *Route {
	return &Route{
		Name:        name,
		Method:      method,
		Pattern:     pattern,
		HandlerFunc: hf,
	}
}
