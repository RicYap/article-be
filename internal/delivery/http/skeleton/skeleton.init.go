package skeleton

import (
	"context"
	"skeleton/internal/entity/skeleton"
)

type SkeletonSvc interface {
	GetAllUser(ctx context.Context) ([]skeleton.User, error)
	GeneratePDF(ctx context.Context) ([]byte, error)
}

type (
	Handler struct {
		skeletonSvc SkeletonSvc
	}
)

func New(is SkeletonSvc) *Handler {
	return &Handler{
		skeletonSvc: is,
	}
}
