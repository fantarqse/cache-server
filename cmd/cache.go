package cachecmd

import (
	"log"
	"net/http"
)

func Run() error {
	// TODO: move to another package
	log.Println("a server is running")
	return http.ListenAndServe(":7777", nil)
}
