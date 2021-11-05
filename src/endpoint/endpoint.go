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
		var solicitud models.Solicitud

		Alumno := function.Alumno_info(db, data.Rut)
		solicitud.Id_alumno = Alumno.Id
		//deberia ser un serial
		solicitud.Id = 821
		//Electivos
		E1 := function.Electivo_info(db, data.Electivo1)
		E2 := function.Electivo_info(db, data.Electivo2)
		E3 := function.Electivo_info(db, data.Electivo2)

		for i := 3; i < 6; i++ {
			var postulacion models.Postulacion
			//esto tampoco deberia estar
			postulacion.Id = i * 456
			//
			postulacion.Aprobado = false
			postulacion.Cantidad = 0
			if i == 3 {
				postulacion.Id_electivo = E1.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_1 = postulacion.Id
			}
			if i == 4 {
				postulacion.Id_electivo = E2.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_2 = postulacion.Id
			}
			if i == 5 {
				postulacion.Id_electivo = E3.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_3 = postulacion.Id
			}
			//	id := function.Insert_postulacion(db, postulacion, logger)
			/* if i == 3 {
				solicitud.Id_Postulacion_1 = id
			}
			if i == 4 {
				solicitud.Id_Postulacion_2 = id
			}
			if i == 5 {
				solicitud.Id_Postulacion_3 = id
			} */
			function.Insert_postulacion(db, postulacion, logger)

		}
		function.Insert_solicitud(db, solicitud, logger)
		logger.Info(Alumno.Nombre)
		c.JSON(200, "OK")
	}
	return
}
