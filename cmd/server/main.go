package main

import (
	"vulnlabz/internal/config"
	"vulnlabz/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	srv.Run()
}
