package sqlgorm

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once = &sync.Once{}

type SQLGorm struct {
	DB *gorm.DB
}

type Config struct {
	GormCfg  gorm.Config
	Username string
	Password string
	HostPort string
	DBName   string
}

func New(cfg Config) SQLGorm {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.HostPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "err gorm open connection"))
	}

	var sqlgorm SQLGorm
	once.Do(func() {
		sqlgorm = SQLGorm{
			DB: db,
		}
	})
	return sqlgorm
}
