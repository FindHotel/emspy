package server

import (
	"context"
	"net/http"

	"github.com/FindHotel/emspy/internal/app/server/handlers"
	"github.com/FindHotel/emspy/internal/app/server/handlers/webhooks"
	"github.com/FindHotel/emspy/internal/app/server/store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv Engine
}

type Engine interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func New(addr string, webhooksStore store.Store) *Server {
	router := gin.Default()
	handlers.RegisterDefaultHandlers(router)

	v1 := router.Group("/v1")
	webhooks.RegisterWebhooks(v1, webhooksStore)

	return &Server{
		srv: &http.Server{Addr: addr, Handler: router},
	}
}
