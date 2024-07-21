package main

import (
	"log"

	initialize "github.com/gildemberg-santos/libratica-go/internal/init"
)

func init() {
	initialize.Run()
}

func main() {
	log.Println("Server is running on http://localhost:8080")
	log.Println("Press Ctrl + C to stop the server")
}
