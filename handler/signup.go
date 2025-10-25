package handler

import (
	"fmt"
	"net/http"

	"github.com/kr0106686/backend/temp"
)

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	email := temp.FormValue(r, "email")
	pass := temp.FormValue(r, "pass")
	name := temp.FormValue(r, "name")
	fmt.Println("가입 요청", email, pass)

	hash := temp.Bcrypt(pass)
	err := h.repo.Signup(name, email, hash)
	if err != nil {
		http.Error(w, "User exists?", 400)
		return
	}

	fmt.Fprint(w, "signup")
}
