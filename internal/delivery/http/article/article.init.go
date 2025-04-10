package article

import (
	"article/internal/entity/article"
	"context"
)

type ArticleSvc interface {
	CreateArticle(ctx context.Context, article article.Post) error
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
