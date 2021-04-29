package gateway

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiv1 "github.com/daangn/guestbook/api/v1"
)

type Guestbook struct {
	apiv1.UnimplementedGuestbookServer
}

func NewGuestbook() *Guestbook {
	return &Guestbook{}
}

func (Guestbook) Create(context.Context, *apiv1.CreateRequest) (*apiv1.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (Guestbook) List(context.Context, *apiv1.ListRequest) (*apiv1.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func (Guestbook) Get(context.Context, *apiv1.GetRequest) (*apiv1.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
