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

	srv := server.New(redis)
	srv.Routes()

	log.Println("a server is running")
	if err := srv.Run(":8080"); err != nil {
		return err
	}

	return nil
}
