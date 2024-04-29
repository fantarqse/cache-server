package main

import (
	"log"

	cache "cache-server/cmd"
)

func main() {
	if err := cache.Run(); err != nil {
		log.Fatalf("failed to start server: %s\n", err.Error())
	}
}
