package skeleton

import (
	"context"
	"skeleton/internal/entity/skeleton"
	"skeleton/pkg/errors"
)

func (d Data) GetAllUser(ctx context.Context) ([]skeleton.User, error) {

	var (
		user  skeleton.User
		users []skeleton.User
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
