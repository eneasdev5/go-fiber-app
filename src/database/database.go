package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func settingsDatabaseConnectRemote() string {
	var (
		DB_USER     string = os.Getenv("DB_USER")
		DB_PASSWORD string = os.Getenv("DB_PASSWORD")
		DB_HOST     string = os.Getenv("DB_HOST")
		DB_PORT     string = os.Getenv("DB_PORT")
		DB_NAME     string = os.Getenv("DB_NAME")
	)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME,
	)
}

func Connect() *sql.DB {
	conn, err := sql.Open("mysql", settingsDatabaseConnectRemote())
	if err != nil {
		panic(err)
	}
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	if err := conn.Ping(); err != nil {
		panic(err)
	}
	return conn
}
