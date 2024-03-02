package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(usr, pass, host, port, database, driver string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", usr, pass, host, port, database)

	db, err := sql.Open(driver, connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
