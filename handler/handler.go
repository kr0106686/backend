package handler

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func New(db *sql.DB) http.Handler {
	h := Handler{db: db}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", h.hello)
	mux.HandleFunc("/users", h.users)

	return mux
}
