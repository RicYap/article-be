package article

import (
	"article/pkg/errors"
	"context"

	"article/internal/entity/article"
)

func (s Service) CreateArticle(ctx context.Context, article article.Posts) error {

	err := s.data.CreateArticle(ctx, article)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetAllUser]")
	}

	return err
}

func (s Service) GetArticleById(ctx context.Context, id int) (article.PostsSlim, error) {

	article, err := s.data.GetArticleById(ctx, id)
	if err != nil {
		return article, errors.Wrap(err, "[SERVICE][GetArticleById]")
	}

	return article, nil
}

func (s Service) GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.PostsSlim, error) {

	offset = (offset - 1) * limit

	article, err := s.data.GetArticlePagination(ctx, limit, offset)
	if err != nil {
		return article, errors.Wrap(err, "[SERVICE][GetArticlePagination]")
	}

	return article, nil
}

func (s Service) UpdateArticle(ctx context.Context, articleBody article.PostsSlim) error {

	err := s.data.UpdateArticle(ctx, articleBody)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateArticle]")
	}

	return nil
}
