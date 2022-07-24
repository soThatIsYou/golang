package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func testRequest(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}
	clientIp := getClientIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientIp: %s", clientIp)
}
