package HomeServices

import (
	"HomeServices/config"
	"HomeServices/core"
	"log"
)

func main() {
	log.Println("Home services bot server v0.1")

	if err := config.InitConfiguration(); err != nil {
		log.Fatalf("Failed to init application: %v", err)
	}

	server, err := core.NewServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


}
