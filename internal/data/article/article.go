package article

import (
	"article/internal/entity/article"
	"article/pkg/errors"
	"context"
)

func (d Data) GetAllUser(ctx context.Context) ([]article.User, error) {

	var (
		user  article.User
		users []article.User
	)

	rows, err := (*d.stmt)[getAllUser].QueryxContext(ctx)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return users, errors.Wrap(err, "[DATA][GetAllUser]")
		}

		users = append(users, user)
	}

	return users, err
}
