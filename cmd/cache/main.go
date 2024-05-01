package main

import (
	"log"

	cachecmd "cache-server/cmd"
)

func main() {
	if err := cachecmd.Run(); err != nil {
		log.Fatalf("failed to start server: %s\n", err.Error())
	}
}
