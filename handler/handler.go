package handler

import (
	"net/http"

	"github.com/kr0106686/backend/repo"
	"github.com/kr0106686/backend/temp"
)

type Handler struct {
	repo *repo.Repo
	jwt  *temp.JWT
}

func New(repo *repo.Repo, jwt *temp.JWT) http.Handler {
	h := Handler{repo: repo, jwt: jwt}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", h.hello)
	mux.HandleFunc("/api/users", h.users)
	mux.HandleFunc("/api/signup", h.signup)
	mux.HandleFunc("/api/login", h.login)
	mux.HandleFunc("/api/profile", h.profile)

	return mux
}
