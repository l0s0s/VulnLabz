package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"vulnlabz/internal/config"
	"vulnlabz/internal/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
	config     *config.Config
}

func New(cfg *config.Config) *Server {
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	routes.SetupRoutes(router)

	httpServer := &http.Server{
		Addr:         cfg.GetAddress(),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	return &Server{
		httpServer: httpServer,
		config:     cfg,
	}
}

func (s *Server) Start() error {
	log.Printf("🚀 VulnLabz server starting on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Run() {
	serverErrors := make(chan error, 1)
	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverErrors:
		if err != nil {
			log.Fatalf("❌ Server error: %v", err)
		}
	case sig := <-shutdown:
		log.Printf("🛑 Received signal: %v. Stopping server...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), s.config.Server.ShutdownTimeout)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Printf("⚠️  Shutdown error: %v", err)
			if err := s.httpServer.Close(); err != nil {
				log.Fatalf("❌ Force close error: %v", err)
			}
		}

		log.Println("✅ Server stopped")
	}
}

func (s *Server) GetAddress() string {
	return fmt.Sprintf("http://%s", s.httpServer.Addr)
}
