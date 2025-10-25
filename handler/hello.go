package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
