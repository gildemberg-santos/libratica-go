package init

import (
	"log"

	"github.com/gildemberg-santos/libratica-go/internal/router"
)

func runGinRouter() {
	log.Println("Initializing Gin router...")
	router.NewRoute().Run()
}
