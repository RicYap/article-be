package article

import (
	"article/internal/entity/article"
	"article/pkg/errors"
	"context"
)

func (d Data) CreateArticle(ctx context.Context, article article.Posts) error {

	err := d.db.
		WithContext(ctx).
		Create(&article).
		Error
	if err != nil {
		return errors.Wrap(err, "[DATA][CreateArticle]")
	}

	return nil
}

func (d Data) GetArticleById(ctx context.Context, id int) (article.Posts, error) {

	var article article.Posts

	err := d.db.
		WithContext(ctx).
		Select("Title, Content, Category, Status").
		Where("Id = ?", id).
		First(&article).
		Error
	if err != nil {
		return article, errors.Wrap(err, "[DATA][GetArticleById]")
	}

	return article, nil
}

func (d Data) GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.Posts, error) {

	var articles []article.Posts

	err := d.db.
		WithContext(ctx).
		Select("Title, Content, Category, Status").
		Limit(limit).
		Offset(offset).
		Find(&articles).
		Error
	if err != nil {
		return articles, errors.Wrap(err, "[DATA][GetArticlePagination]")
	}

	return articles, nil
}


