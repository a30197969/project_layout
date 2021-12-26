package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 数据库驱动
	"project_layout/internal/conf"
)

type DataSource interface {
	Conn(driver string, source string)
}

type Mysql struct {
	db *sql.DB
}

func (m *Mysql) Conn(driver string, source string) {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}
	m.db = db
}

func NewMySQL(cfg *conf.Config) DataSource {
	m := &Mysql{}
	m.Conn(cfg.Database.Driver, cfg.Database.Source)
	return m
}
