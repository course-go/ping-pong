package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = "8080"
)

type Response struct {
	Message  string    `json:"message,omitempty"`
	Instance string    `json:"instance,omitempty"`
	Time     time.Time `json:"time,omitempty"`
}

func main() {
	instance := os.Getenv("INSTANCE")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message:  "pong",
			Instance: instance,
			Time:     time.Now(),
		}

		bytes, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		slog.Debug("handled HTTP request")
		w.Write(bytes)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	hostname := fmt.Sprintf(":%s", port)
	slog.Info("starting server", "hostname", hostname)
	err := http.ListenAndServe(hostname, mux)
	slog.Error("failed running server", "error", err)
}
