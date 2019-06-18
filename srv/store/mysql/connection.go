package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")))

	// if error occur, log it for now
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return db, nil
}
