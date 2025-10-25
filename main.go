package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kr0106686/backend/handler"
	"github.com/kr0106686/backend/temp"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = ".env"
	}

	err := godotenv.Load(env)
	if err != nil {
		log.Fatal(err)
	}

	db, err := temp.PQ()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✅ DB Connected!")

	h := handler.New(db)

	http.ListenAndServe(":3000", h)
}
