package handler

import (
	"fmt"
	"net/http"

	"github.com/kr0106686/backend/temp"
)

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	email := temp.FormValue(r, "email")
	pass := temp.FormValue(r, "pass")
	fmt.Println("login요청", email, pass)

	if email == "" || pass == "" {
		http.Error(w, "missing required value", http.StatusUnauthorized)
		return
	}

	hash, err := h.repo.Login(email)
	if err != nil {
		http.Error(w, "user not exists", http.StatusUnauthorized)
		return
	}

	if err := temp.CompareHashAndPassword(hash, pass); err != nil {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	token := h.jwt.Issue(email)
	if token == "" {
		http.Error(w, "token 생성 실패", 400)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	fmt.Fprint(w, "login")
}
