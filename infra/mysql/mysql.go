package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	Username string
	Password string
	Host     string
	Port     string
	DBName   string

	// Default is unlimited
	// set to lest than equal to 0 means make it unlimited or default setting,
	// more open connection means less time taken to perform query
	MaxOpenConn int

	// MaxIdleConn default is 2
	// set to lest than equal to 0 means not allow any idle connection,
	// more idle connection in the pool will improve performance,
	// since no need to establish connection from scratch)
	// by set idle connection to 0, a new connection has to be created from scratch for each operation
	// ! should be <= MaxOpenConn
	MaxIdleConn int

	// MaxLifetime set max length of time that a connection can be reused for.
	// Setting to 0 means that there is no maximum lifetime and
	// the connection is reused forever (which is the default behavior)
	// the shorter lifetime result in more memory useage
	// since it will kill the connection and recreate it
	MaxLifetime time.Duration
}

// New create new MySQL instance
func New(cfg Config) MySQL {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		panic(err.Error())
	}

	if cfg.MaxOpenConn != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	if cfg.MaxIdleConn != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConn)
	}

	if err = db.Ping(); err != nil {
		logrus.Fatal(errors.Wrap(err, "ping mysql"))
	}

	logrus.Infof("%-7s %s", "MySQL", "✅")

	return db
}
