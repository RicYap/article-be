package article

import (
	"context"
	"article/internal/entity/article"
)

type Data interface {
	GetAllUser(ctx context.Context) ([]article.User, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
