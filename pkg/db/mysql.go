package db

import (
	"context"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	dsn string
	cfg *gorm.Config
	db  *gorm.DB
}

type Option func(*MySQL)

func SetDSN(dsn string) Option {

	return func(ms *MySQL) {
		ms.dsn = dsn
	}
}

func WithConfig(cfg *gorm.Config) Option {

	return func(ms *MySQL) {
		ms.cfg = cfg
	}
}
func NewMySQL(options ...Option) *MySQL {

	mySQL := new(MySQL)

	for _, v := range options {
		v(mySQL)
	}

	return mySQL

}
func (m *MySQL) Start(ctx context.Context) error {
	if m.dsn == "" || m.cfg == nil {
		return errors.New("mysql dsn or config need init")
	}
	var err error
	if m.db, err = gorm.Open(mysql.Open(m.dsn), m.cfg); err != nil {
		return err
	}
	return nil
}

func (m *MySQL) GetString() string {
	return "MySQLDB"
}

func (m *MySQL) Stop() error {
	return nil
}
