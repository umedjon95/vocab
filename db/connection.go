package db

import (
	"database/sql"
	"fmt"
	"vocab/config"
)

var (
	db  *sql.DB
	err error
)

//Connect is a function to connect to DB
func Connect() error {
	dbConf := config.Peek().Database
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Addr, dbConf.User, dbConf.Pass, dbConf.DBName)
	db, err = sql.Open("postgres", psqlInfo)

	return err
}

//Close is a function to disconnect from DB
func Close() error {
	return db.Close()
}

//Ping is a function to check the connection with DB
func Ping() (err error) {
	err = db.Ping()
	return
}
