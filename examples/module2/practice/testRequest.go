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
	// 02.将requst中的header 设置到 reponse中
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}
	// 04.记录日志并输出
	clientip := getClientIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientip)
}
