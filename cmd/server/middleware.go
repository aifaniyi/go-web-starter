package main

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/aifaniyi/sample/pkg/logger"
)

type key int

const (
	requestID key = iota
)

func (s *server) loggingMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		writer := responseWriter{ResponseWriter: w}
		h.ServeHTTP(w, r)

		duration := time.Since(start)
		referer := r.Referer()
		if referer == "" {
			referer = "-"
		}

		ua := r.UserAgent()
		if ua == "" {
			ua = "-"
		}

		id := "-"
		if reqID := r.Context().Value(requestID); reqID != nil {
			id = reqID.(string)
		}

		// format: req-id, http-method  requestURI  http-protocol  useragent  remote-addr
		// referer  request-length  time-taken response-status response-length
		logger.Info.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%v\t%d\t%d",
			id, r.Method, r.RequestURI, r.Proto, ua,
			r.RemoteAddr, referer, r.ContentLength, duration,
			writer.Status, writer.Length,
		)
	})
}

// responseWriter : response writer with
// response status and response size information
type responseWriter struct {
	http.ResponseWriter
	Status int
	Length int
}

// WriteHeader : writes response status
func (w *responseWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

// Write : writes response data to writer
func (w *responseWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.Length += n
	return n, err
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}
