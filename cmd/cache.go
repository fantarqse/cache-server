package cachecmd

import (
	"cache-server/internal/cache"
	"cache-server/internal/config"
	"cache-server/internal/server"
)

func Run() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	redis := cache.New(cfg.Redis)

	srv := server.New(redis)
	srv.Routes()

	if err := srv.Run(cfg.Server.Port); err != nil {
		return err
	}

	return nil
}
