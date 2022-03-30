package account

import (
	"context"

	mysqld "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/mysql"
)

type Repository interface {
	FindByUsername(ctx context.Context, username string) (Account, error)
	CreateAccount(ctx context.Context, v Account) error
}

type repository struct {
	sql mysql.MySQL
}

func NewRepository(sql mysql.MySQL) Repository {
	return &repository{
		sql: sql,
	}
}

func (r *repository) CreateAccount(ctx context.Context, v Account) error {
	query := `INSERT INTO user(
		id,
		username,
		password,
		name,
		email,
		phone,
		is_active,
		is_superuser,
		created_at,
		updated_at,
		last_login
	) VALUES(?,?,?,?,?,?,?,?,?,?,?)`

	_, err := r.sql.ExecContext(ctx, query,
		v.ID,
		v.Username,
		v.password,
		v.Name,
		v.Email,
		v.Phone,
		v.IsActive,
		v.IsSuperuser,
		v.CreatedAt,
		v.UpdatedAt,
		v.LastLogin,
	)
	if err != nil {
		if mysqlErr, ok := err.(*mysqld.MySQLError); ok {
			if mysqlErr.Number == mysql.ErrDuplicateKey {
				return errors.Wrap(err, "duplicate key")
			}
		}
		return errors.Wrap(err, "exec query")
	}

	return nil
}

func (r *repository) FindByUsername(ctx context.Context, username string) (acc Account, err error) {
	query := `SELECT 
		id,
		username,
		password,
		name,
		email,
		phone,
		is_active,
		is_superuser,
		created_at,
		updated_at,
		last_login 
	FROM user 
	WHERE username = ?`

	row := r.sql.QueryRowContext(ctx, query, username)
	if err := row.Scan(
		&acc.ID,
		&acc.Username,
		&acc.password,
		&acc.Name,
		&acc.Email,
		&acc.Phone,
		&acc.IsActive,
		&acc.IsSuperuser,
		&acc.CreatedAt,
		&acc.UpdatedAt,
		&acc.LastLogin,
	); err != nil {
		return acc, errors.Wrap(err, "scan row")
	}

	return acc, err
}
