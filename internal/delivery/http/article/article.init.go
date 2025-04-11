package article

import (
	"article/internal/entity/article"
	"context"
)

type ArticleSvc interface {
	CreateArticle(ctx context.Context, article article.Posts) error
	GetArticleById(ctx context.Context, id int) (article.Posts, error)
	GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.Posts, error)
}

type (
	Handler struct {
		articleSvc ArticleSvc
	}
)

func New(is ArticleSvc) *Handler {
	return &Handler{
		articleSvc: is,
	}
}
