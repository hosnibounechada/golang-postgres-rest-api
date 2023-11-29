package main

import (
	"log"

	"github.com/hosnibounechada/go-api/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app, err := internal.NewApp()

	if err != nil {
		log.Fatalf("Failed to create the app: %v", err)
	}

	app.Run()
}
