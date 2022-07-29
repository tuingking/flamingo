package product

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/infra/qbuilder"
	"github.com/tuingking/flamingo/infra/sqlgorm"
	"github.com/tuingking/flamingo/internal/app"
)

type Repository interface {
	FindAll(ctx context.Context, p GetProductParam) ([]Product, app.Pagination, error)
	Create(ctx context.Context, v Product) (Product, error)
}

type repository struct {
	mysql  mysql.MySQL
	gormdb sqlgorm.SQLGorm
}

func NewRepository(mysql mysql.MySQL, gormdb sqlgorm.SQLGorm) Repository {
	return &repository{
		mysql:  mysql,
		gormdb: gormdb,
	}
}

func (r *repository) FindAll(ctx context.Context, p GetProductParam) (res []Product, pagination app.Pagination, err error) {
	// set default pagination
	p.Page, p.Limit = qbuilder.ValidatePageAndLimit(p.Page, p.Limit)

	query := `SELECT id, name FROM product`

	wc, args, err := qbuilder.New().Build(&p)
	if err != nil {
		return res, pagination, errors.Wrap(err, "build qbuilder")
	}

	rows, err := r.mysql.QueryContext(ctx, query+wc, args...)
	if err != nil {
		return res, pagination, errors.Wrap(err, "exec query")
	}
	defer rows.Close()

	for rows.Next() {
		var v Product
		if err := rows.Scan(
			&v.ID,
			&v.Name,
		); err != nil {
			return res, pagination, errors.Wrap(err, "row scan")
		}

		res = append(res, v)
	}

	// pagination
	pagination.CurrentPage = p.Page
	pagination.PageSize = p.Limit
	pagination.TotalElement = int64(len(res))

	return res, pagination, nil
}

func (r *repository) Create(ctx context.Context, v Product) (Product, error) {
	query := `INSERT INTO product(name, price) VALUES(?,?)`

	res, err := r.mysql.ExecContext(ctx, query, v.Name, v.BuyPrice)
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
