package app

import (
	"log"
	"net/http"

	"github.com/kr0106686/backend/handler"
	"github.com/kr0106686/backend/repo"
	"github.com/kr0106686/backend/temp"
)

func New() *http.Server {
	db, err := temp.PQ()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✅ DB Connected!")

	r := repo.New(db)
	j := temp.New()

	h := handler.New(r, j)

	return &http.Server{
		Addr:    ":3000",
		Handler: h,
	}
}
