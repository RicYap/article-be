package article

import (
	"article/internal/entity/article"
	"context"
)

type ArticleSvc interface {
	GetAllUser(ctx context.Context) ([]article.User, error)
	GeneratePDF(ctx context.Context) ([]byte, error)
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
