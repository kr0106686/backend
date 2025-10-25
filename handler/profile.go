package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) profile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "No auth", http.StatusUnauthorized)
		return
	}

	tokenStr := cookie.Value
	email := h.jwt.Parse(tokenStr)
	if email == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "email : %s", email)
}
