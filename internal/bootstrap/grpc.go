package bootstrap

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	apiv1 "github.com/daangn/guestbook/api/v1"
	"github.com/daangn/guestbook/internal/config"
	"github.com/daangn/guestbook/internal/gateway"
)

type Guestbook struct {
	config *config.Config
	server *gateway.Guestbook
}

func NewGuestbook(config *config.Config, server *gateway.Guestbook) *Guestbook {
	return &Guestbook{
		config: config,
		server: server,
	}
}

func (g Guestbook) Serve() error {
	listener, err := net.Listen("tcp", g.config.Listen)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	apiv1.RegisterGuestbookServer(srv, g.server)
	reflection.Register(srv)

	return srv.Serve(listener)
}
