// main.go

package main

import (
	"database/sql"
	"electivos-ucn/src/database"
	"electivos-ucn/src/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func setupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	api := r.Group("/ucn")
	{
		public := api.Group("/public")
		{
			public.POST("/login")
			//public.POST("/data", controllers.TestEndPoint)
		}

	}
	return r
}

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&utils.LogFormat{})
	logger := logrus.WithFields(nil)
	logger.Info("Initializing app...")

	db := database.DBConnection(logger)
	r := setupRouter(db)
	port := os.Getenv("PORT")
	r.Run(port)
}
