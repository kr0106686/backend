package main

import (
	"net/http"

	"github.com/kr0106686/backend/handler"
)

func main() {
	http.HandleFunc("/api/hello", handler.Hello)
	http.ListenAndServe(":3000", nil)
}
