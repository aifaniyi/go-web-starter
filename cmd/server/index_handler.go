package main

import (
	"net/http"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpResponse(w, "Hello World!", http.StatusOK)
	}
}
