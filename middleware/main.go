package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		h.ServeHTTP(w, r)
		d := time.Now().Sub(s).Milliseconds()
		log.Printf("end %s(%d ms)\n", time.Now().Format(time.RFC3339), d)
	})
}

func VersionAdd(v string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("App-Version", v)
			next.ServeHTTP(w, r)
		})

	}
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				jsonBody, _ := json.Marshal(map[string]string{
					"error": fmt.Sprintf("%v", err),
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to log request body", zap.Error(err))
			http.Error(w, "Failed to get request body", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}
