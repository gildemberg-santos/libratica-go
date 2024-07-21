package init

import (
	"log"

	"github.com/joho/godotenv"
)

func runEnv() {
	log.Println("Loading environment variables...")
	godotenv.Load()
}
