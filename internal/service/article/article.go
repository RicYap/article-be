package article

import (
	"article/pkg/errors"
	"context"

	"article/internal/entity/article"
)

func (s Service) CreateArticle(ctx context.Context, article article.Post) error {

	err := s.data.CreateArticle(ctx, article)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUser]")
	}

	return err
}
