package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test/request", testRequest)
	mux.HandleFunc("/healthz", checkHealth)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}
