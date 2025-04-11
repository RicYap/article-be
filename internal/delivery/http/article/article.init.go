package article

import (
	"article/internal/entity/article"
	"context"
)

type ArticleSvc interface {
	CreateArticle(ctx context.Context, article article.Posts) error
	GetArticleById(ctx context.Context, id int) (article.PostsResponse, error)
	GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.PostsResponse, error)
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
