package main

import (
	"net/http"

	"github.com/kr0106686/backend/handler"
)

func main() {
	http.HandleFunc("/hello", handler.Hello)
	http.ListenAndServe(":3000", nil)
}
