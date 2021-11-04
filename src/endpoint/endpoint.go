package endpoint

import (
	"database/sql"
	"electivos-ucn/src/function"
	"electivos-ucn/src/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Formulario(c *gin.Context, db *sql.DB, logger *logrus.Entry) (data models.Formulario) {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	} else {
		//TODO EL CODIGO PARA BD
		Alumno := function.Alumno_info(db, data.Rut)
		logger.Info(Alumno.Nombre)
		c.JSON(200, "OK")
	}
	return
}
