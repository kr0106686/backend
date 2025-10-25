package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) users(w http.ResponseWriter, r *http.Request) {

	var count int
	err := h.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	fmt.Fprint(w, count)
}
