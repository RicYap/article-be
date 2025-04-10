package article

import (
	"context"
	"article/internal/entity/article"
)

type Data interface {
	CreateArticle(ctx context.Context, article article.Post) error
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
