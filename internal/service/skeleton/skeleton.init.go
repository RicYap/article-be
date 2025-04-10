package skeleton

import (
	"context"
	"skeleton/internal/entity/skeleton"
)

type Data interface {
	GetAllUser(ctx context.Context) ([]skeleton.User, error)
}

type Service struct {
	data    Data
}

func New(data Data) Service {
	return Service{
		data:    data,
	}
}
