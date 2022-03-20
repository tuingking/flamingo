package product

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/mysql"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Product, error)
}

type repository struct {
	mysql mysql.MySQL
}

func NewRepository(mysql mysql.MySQL) Repository {
	return &repository{
		mysql: mysql,
	}
}

func (r *repository) FindAll(ctx context.Context) ([]Product, error) {
	var res []Product

	query := `SELECT id, name FROM product`

	rows, err := r.mysql.QueryContext(ctx, query)
	if err != nil {
		return res, errors.Wrap(err, "exec query")
	}

	for rows.Next() {
		var v Product
		if err := rows.Scan(
			&v.ID,
			&v.Name,
		); err != nil {
			return res, errors.Wrap(err, "row scan")
		}

		res = append(res, v)
	}

	return res, nil
}
