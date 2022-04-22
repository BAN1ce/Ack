package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string
	srv  *http.Server
}

type Option func(*Server)

func SetAddr(addr string) Option {

	return func(s *Server) {

		s.addr = addr
	}
}

func NewServer(options ...Option) *Server {
	s := new(Server)

	for _, v := range options {
		v(s)
	}
	return s
}

// s *Server provider.IProvider

func (s *Server) Start(ctx context.Context) error {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	if s.addr == "" {
		s.addr = ":8080"
	}
	s.srv = &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	return nil
}
func (s *Server) GetString() string {

	return "ginServer"
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
