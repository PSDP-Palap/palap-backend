package main

import (
	"log"
	"palap_backend/routers"
)

func main() {
	r := routers.InitializeRouter()

	if err := r.Run(); err != nil {
		log.Printf("failed to run server: %v", err)
	}
}
