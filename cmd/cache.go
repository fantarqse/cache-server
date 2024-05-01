package cachecmd

import (
	"log"

	"cache-server/internal/cache"
	"cache-server/internal/server"
)

// TODO: move to the config file
const addr string = "localhost:6379"

func Run() error {
	log.Println("a redis is running")
	redis := cache.New(addr)

	log.Println("a server is running")
	if err := server.New(redis).Run(":8080"); err != nil { // TODO: move the port to the config
		return err
	}

	return nil
}
