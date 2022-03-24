package product

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/mysql"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Product, error)
	Create(ctx context.Context, v Product) (Product, error)
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
	defer rows.Close()

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

func (r *repository) Create(ctx context.Context, v Product) (Product, error) {
	// dbStat := r.mysql.Stats()
	// now := time.Now()
	// defer func() {
	// 	logrus.WithFields(logrus.Fields{
	// 		"MaxOpenConnections": dbStat.MaxOpenConnections,
	// 		"OpenConnections":    dbStat.OpenConnections,
	// 		"InUse":              dbStat.InUse,
	// 		"Idle":               dbStat.Idle,
	// 		"Elapsed":            time.Since(now),
	// 	}).Info(v.Name)
	// }()

	query := `INSERT INTO product(name, price) VALUES(?,?)`

	res, err := r.mysql.ExecContext(ctx, query, v.Name, v.Price)
	if err != nil {
		return v, errors.Wrap(err, "exec query")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return v, errors.Wrap(err, "get last insert ID")
	}
	v.ID = id

	return v, nil
}
