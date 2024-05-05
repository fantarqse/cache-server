package cachecmd

import (
	"cache-server/internal/config"
	"cache-server/internal/server"
	"cache-server/internal/storage"
)

func Run() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	redis := storage.New(cfg.Redis)

	if err := server.New(redis).Run(cfg.Server.Port); err != nil {
		return err
	}

	return nil
}
