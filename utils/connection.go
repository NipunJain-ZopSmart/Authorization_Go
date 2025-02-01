package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connection(user, password, host, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)

	//ATTEMPT TO CONNECT TO DB
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to Db!!")
	return db, nil
}
