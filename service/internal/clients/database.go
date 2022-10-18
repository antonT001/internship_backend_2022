package clients

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"user_balance/service/internal/logger"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type DataBaseMethods interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	DeleteIn(query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
}

type DataBase interface {
	DataBaseMethods
	NewTransaction() (Transaction, error)
}

type dataBase struct {
	db     *sqlx.DB
	logger logger.Logger
}

type Transaction interface {
	DataBaseMethods
	Rollback() error
	Commit() error
}

type sqlxTransaction struct {
	*sqlx.Tx
}

func New(logger logger.Logger) DataBase {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWD")
	dbname := os.Getenv("MYSQL_DBNAME")

	source := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		user, pass, host, port, dbname,
	)

	conn, err := sqlx.Connect("mysql", source)
	if err != nil {
		//logger.Panic(err)
	} else {
		//conn.SetConnMaxIdleTime(5 * time.Second)
		conn.SetConnMaxLifetime(60 * time.Second)
		conn.SetMaxIdleConns(10)
		conn.SetMaxOpenConns(10)
	}
	return &dataBase{
		db:     conn,
		logger: logger,
	}
}

func (d *dataBase) NewTransaction() (Transaction, error) {
	tx, _ := d.db.Beginx()
	return &sqlxTransaction{tx}, nil
}

func (d *dataBase) Select(dest interface{}, query string, args ...interface{}) error {
	return d.db.Select(dest, query, args...)
}

func (d *dataBase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d *dataBase) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return d.db.NamedExec(query, arg)
}

func (d *dataBase) DeleteIn(query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = d.db.Exec(query, inArgs...)
	return err
}

func (d *dataBase) Get(dest interface{}, query string, args ...interface{}) error {
	return d.db.Get(dest, query, args...)
}

func (d *sqlxTransaction) DeleteIn(query string, args ...interface{}) error {
	query, inArgs, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	_, err = d.Exec(query, inArgs...)
	return err
}
