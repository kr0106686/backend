package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) users(w http.ResponseWriter, r *http.Request) {
	count, err := h.repo.Users()
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	fmt.Fprint(w, count)
}
