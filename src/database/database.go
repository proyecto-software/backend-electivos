// database/database.go

package database

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// GlobalDB a global db object will be used across different packages
var GlobalDB *gorm.DB

func DBConnection(logger *logrus.Entry) *sql.DB {
	db, err := sql.Open("postgres", "host=ec2-52-72-252-211.compute-1.amazonaws.com port=5432 user=spdkbddxynexbe password=2244405d1de4d94e500c668868571ea324040709be8beb1ae2fc1769e311a91d dbname=d4qsaefpehhcjp sslmode=enable")

	if err != nil {
		log.Error("Error when try to connect db")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	logger.Info("Successfully connected!")

	return db
}
