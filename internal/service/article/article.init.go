package article

import (
	"article/internal/entity/article"
	"context"
)

type Data interface {
	CreateArticle(ctx context.Context, article article.Posts) error
	GetArticleById(ctx context.Context, id int) (article.PostsResponse, error)
	GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.PostsResponse, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
