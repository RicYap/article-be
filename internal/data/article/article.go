package article

import (
	"article/internal/entity/article"
	"article/pkg/errors"
	"context"
	"time"
)

func (d Data) CreateArticle(ctx context.Context, article article.Posts) error {

	article.CreatedDate = time.Now().Local()
	article.UpdatedDate = time.Now().Local()

	err := d.db.
		WithContext(ctx).
		Create(&article).
		Error
	if err != nil {
		return errors.Wrap(err, "[DATA][CreateArticle]")
	}

	return nil
}

func (d Data) GetArticleById(ctx context.Context, id int) (article.PostsSlim, error) {

	var articleResult article.PostsSlim

	err := d.db.
		WithContext(ctx).
		Model(&article.Posts{}).
		Select("Title, Content, Category, Status").
		Where("Id = ?", id).
		First(&articleResult).
		Error
	if err != nil {
		return articleResult, errors.Wrap(err, "[DATA][GetArticleById]")
	}

	return articleResult, nil
}

func (d Data) GetArticlePagination(ctx context.Context, limit int, offset int) ([]article.PostsSlim, error) {

	var articles []article.PostsSlim

	err := d.db.
		WithContext(ctx).
		Model(&article.Posts{}).
		Select("Id, Title, Content, Category, Status").
		Limit(limit).
		Offset(offset).
		Find(&articles).
		Error
	if err != nil {
		return articles, errors.Wrap(err, "[DATA][GetArticlePagination]")
	}

	return articles, nil
}

func (d Data) GetArticlePaginationByStatus(ctx context.Context, limit int, offset int, status string) ([]article.PostsSlim, error) {

	var articles []article.PostsSlim

	err := d.db.
		WithContext(ctx).
		Model(&article.Posts{}).
		Select("Id, Title, Content, Category, Status").
		Where("status LIKE ?", "%"+status+"%").
		Limit(limit).
		Offset(offset).
		Find(&articles).
		Error
	if err != nil {
		return articles, errors.Wrap(err, "[DATA][GetArticlePaginationByStatus]")
	}

	return articles, nil
}

func (d Data) UpdateArticle(ctx context.Context, articleBody article.PostsSlim) error {

	err := d.db.
		WithContext(ctx).
		Model(&article.Posts{}).
		Where("Id = ?", articleBody.ID).
		Select("Title", "Content", "Category", "Status", "Updated_date").
		Updates(map[string]interface{}{
			"Title":    articleBody.Title,
			"Content":  articleBody.Content,
			"Category": articleBody.Category,
			"Status":   articleBody.Status,
			"Updated_date": time.Now().Local(),
		}).Error

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateArticle]")
	}

	return nil
}

func (d Data) DeleteArticle(ctx context.Context, id int) error {

	err := d.db.
		WithContext(ctx).
		Delete(&article.Posts{}, id).
		Error
	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteArticle]")
	}

	return nil
}
