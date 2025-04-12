package article

import (
	"article/internal/entity/article"
	"context"
)

type ArticleSvc interface {
	CreateArticle(ctx context.Context, article article.Posts) error
	GetArticleById(ctx context.Context, id int) (article.PostsSlim, error)
	GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.PostsSlim, error)
	GetArticlePaginationByStatus(ctx context.Context, limit int, offset int, status string) ([]article.PostsSlim, error)
	UpdateArticle(ctx context.Context, articleBody article.PostsSlim) error
	DeleteArticle(ctx context.Context, id int) error
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
