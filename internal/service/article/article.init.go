package article

import (
	"article/internal/entity/article"
	"context"
)

type Data interface {
	CreateArticle(ctx context.Context, article article.Posts) error
	GetArticleById(ctx context.Context, id int) (article.PostsSlim, error)
	GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.PostsSlim, error)
	UpdateArticle(ctx context.Context, articleBody article.PostsSlim) error
	DeleteArticle(ctx context.Context, id int) error
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
