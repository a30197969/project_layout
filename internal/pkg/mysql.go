package pkg

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 数据库驱动
)

type DataSource interface {

}

type Mysql struct {
	db *sql.DB
}

//func NewMySQL(conf *DBConfig) DataSource {
//	sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/fn_bird")
//	return &mysql{db:}
//}
