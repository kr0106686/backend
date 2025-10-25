package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kr0106686/backend/app"
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

	a := app.New()
	a.ListenAndServe()
}
