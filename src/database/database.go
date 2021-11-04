// database/database.go

package database

import (
	"database/sql"
	"os"

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
	db, err := sql.Open("postgres", os.Getenv("SQL_CONNECTION_STRING"))

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
