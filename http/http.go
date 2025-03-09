package http

import (
	"fmt"
	"net/http"

	nginxcachepurger "github.com/dkvz/nginx-cache-purger"
)

type Server struct {
	requester *nginxcachepurger.Requester
}

func NewServer(config *nginxcachepurger.Config) *Server {
	requester := nginxcachepurger.NewRequester(config)
	requester.Start()
	return &Server{
		requester,
	}
}

func (s *Server) ListenAndServe(port uint) error {

	return http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", port), nil)
}
