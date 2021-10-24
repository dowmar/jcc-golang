package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	username string = "root"
// 	password string = "domar123"
// 	database string = "akademik"
// )

// var (
// 	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
// )

var (
	username string = os.Getenv("DB_USERNAME")
	password string = os.Getenv("DB_PASSWORD")
	database string = os.Getenv("DB_DATABASE")
	host     string = os.Getenv("DB_HOST")
)

// const (
// 	localUsername string = "root"
// 	localPassword string = "domar123"
// 	localDatabase string = "db_ecomm"
// )

var (
	dsn = fmt.Sprintf("%v:%v@%v/%v", username, password, host, database)
)

// HubToMySQL
func MySQL() (*sql.DB, error) {
	// var dsn string

	// if username == " " {
	// 	dsn = fmt.Sprintf("%v:%v@%v/%v", username, password, host, database)
	// } else {
	// 	dsn = fmt.Sprintf("%v:%v@/%v", localUsername, localPassword, localDatabase)
	// }

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
