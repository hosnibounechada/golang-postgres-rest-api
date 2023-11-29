package main

import "github.com/hosnibounechada/go-api/internal/auth"

func main() {
	app, err := auth.NewApp() // Create a new instance of your application.

	if err != nil {
		return
	}
	// You can perform additional setup or configuration here if needed.

	// Start the server.
	app.Run()
}
