package main

import (
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	slog.Info("Server listening to :8080")
	http.ListenAndServe(":8080", mux)
}
