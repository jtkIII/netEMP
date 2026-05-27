package main

import (
	"log"

	"netEMP/internal/server"
)

func main() {
	srv := server.New()
	log.Println("starting server on :8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
