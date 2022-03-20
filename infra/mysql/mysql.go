package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
	Ping() error
	Close() error
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	SetConnMaxIdleTime(d time.Duration)
	Stats() sql.DBStats
}

type Config struct {
	Username    string
	Password    string
	Host        string
	Port        string
	DBName      string
	MaxOpenConn int
	MaxIdleConn int
}

// New create new MySQL instance
func New(cfg Config) MySQL {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		panic(err.Error())
	}

	if cfg.MaxOpenConn != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	if cfg.MaxIdleConn != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConn)
	}

	return db
}
