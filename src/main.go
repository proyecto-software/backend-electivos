// main.go

package main

import (
	"database/sql"
	"electivos-ucn/src/database"
	"electivos-ucn/src/endpoint"
	"electivos-ucn/src/function"
	"electivos-ucn/src/middlewares"
	"electivos-ucn/src/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//"electivos-ucn/src/function"
func setupRouter(db *sql.DB, logger *logrus.Entry) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	api := r.Group("/ucn")
	{
		api.POST("/formulario", func(c *gin.Context) {
			endpoint.Formulario(c, db, logger)
		})
		api.POST("/dashboard", func(c *gin.Context) {
			endpoint.Formulario(c, db, logger)
		})
		api.GET("/electivos", func(c *gin.Context) {
			endpoint.GetElectivos(c, db, logger)
		})
		api.GET("/carreras", func(c *gin.Context) {
			endpoint.GetCarrera(c, db, logger)
		})
		api.GET("/email", func(c *gin.Context) {
			function.SendEmail2("nicolas.garcia@alumnos.ucn.cl")
			function.SendEmail2("ssp013@alumnos.ucn.cl")
			function.SendEmail2("jose.flores02@alumnos.ucn.cl")
			function.SendEmail2("dionisio.olivares@alumnos.ucn.cl")
		})
		api.POST("/rut", func(c *gin.Context) {
			endpoint.Rut(c, db, logger)
		})
		api.POST("/correo", func(c *gin.Context) {
			endpoint.Correo(c, db, logger)
		})

		api.GET("/solicitud", func(c *gin.Context) {
			endpoint.GetSolicitud(c, db, logger)
		})
		//indicador
		/* api.GET("/Indicador", func(c *gin.Context) {
		}) */
		//para ver la info de un alumno en particular
		/* api.GET("/{rut}", func(c *gin.Context) {
			function.InformeCurricular(c, db, logger)
		}) */

		api.GET("/InformeCurricular", func(c *gin.Context) {
			endpoint.InformeCurricular(c, db, logger)
		})
		api.POST("/AprobarPostulacion", func(c *gin.Context) {
			endpoint.EstadoPostulacion(c, db, logger)
		})
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
