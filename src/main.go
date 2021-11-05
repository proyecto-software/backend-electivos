// main.go

package main

import (
	"database/sql"
	"electivos-ucn/src/database"
	"electivos-ucn/src/endpoint"
	"electivos-ucn/src/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//"electivos-ucn/src/function"
func setupRouter(db *sql.DB, logger *logrus.Entry) *gin.Engine {
	r := gin.Default()
	api := r.Group("/ucn")
	{
		api.POST("/formulario", func(c *gin.Context) {
			endpoint.Formulario(c, db, logger)
		})
		api.POST("/dashboard", func(c *gin.Context) {
			endpoint.Formulario(c, db, logger)
		})
		/*  api.GET("/alumnos", func(c *gin.Context) {
			function.getAlumnos(c,db,logger)
		})  */
		/* api.VIEW("/alumnos/{rut}", func(c *gin.Context) {
			function.createIndicador(c, db, logger)
		}) */
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
	r := setupRouter(db, logger)
	port := os.Getenv("PORT")
	r.Run(port)
}
