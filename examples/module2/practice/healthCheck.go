package main

import (
	"fmt"
	"net/http"
)

// 05.健康检查的路由
func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
