package article

import (
	"article/internal/entity/article"
	"article/pkg/errors"
	"context"
)

func (d Data) CreateArticle(ctx context.Context, article article.Post) error {

	err := d.db.
		WithContext(ctx).
		Create(&article)
	if err != nil {
		return errors.Wrap(err.Error, "[DATA][CreateArticle]")
	}

	return nil
}

