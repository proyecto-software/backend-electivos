package endpoint

import (
	"database/sql"
	"electivos-ucn/src/function"
	"electivos-ucn/src/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//"github.com/gorilla/mux"
//ejecutar -> go mod tidy
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
		var E1, E2, E3 models.Electivo
		if data.Electivo1 != "" {
			E1 = function.Electivo_info(db, data.Electivo1)
		}
		if data.Electivo2 != "" {
			E2 = function.Electivo_info(db, data.Electivo2)
		}
		if data.Electivo3 != "" {
			E3 = function.Electivo_info(db, data.Electivo3)
		}

		for i := 3; i < 6; i++ {
			var postulacion models.Postulacion
			//esto tampoco deberia estar
			postulacion.Id = i * 456
			//
			postulacion.Aprobado = false
			postulacion.Cantidad = data.Cantidad
			if data.Cantidad == 1 {
				postulacion.Id_electivo = E1.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_1 = postulacion.Id
			}
			if data.Cantidad == 2 {
				postulacion.Id_electivo = E1.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_1 = postulacion.Id
				postulacion.Id_electivo = E2.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_2 = postulacion.Id
			}
			if data.Cantidad == 3 {
				postulacion.Id_electivo = E1.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_1 = postulacion.Id
				postulacion.Id_electivo = E2.Id
				// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
				solicitud.Id_Postulacion_2 = postulacion.Id
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

func InformeCurricular(c *gin.Context, db *sql.DB, logger *logrus.Entry) (data models.Alumno) {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	} else {
		var_rut := c.DefaultQuery("rut", "")
		Alumno := function.Alumno_info(db, var_rut)
		logger.Infof(Alumno.Nombre)
		/*
			vars := mux.Vars(r) //guarda el rut
			alumnoRut := vars["rut"] //tiene que llamarse igual que en el router

			function.Alumno_info(db, alumnoRut)
			//function.getAlumno()
		*/
	}
	return
}

func GetElectivos(c *gin.Context, db *sql.DB, logger *logrus.Entry) {

	carrera := c.DefaultQuery("carrera", "")
	if carrera == "" {
		data := function.All_electivo_info(db)
		c.JSON(200, data)
	}

}

func GetCarrera(c *gin.Context, db *sql.DB, logger *logrus.Entry) {
	carrera := c.DefaultQuery("carrera", "")
	if carrera == "" {
		data := function.All_carrera_info(db)
		c.JSON(200, data)
	}

}

func Rut(c *gin.Context, db *sql.DB, logger *logrus.Entry) {
	var data models.Formulario
	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	} else {
		valid := function.Validator(data.Rut, logger, c)
		if valid {
			c.JSON(200, "OK")
		}
	}

}

func Correo(c *gin.Context, db *sql.DB, logger *logrus.Entry) {
	var data models.Formulario
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	} else {
		valid := function.ValidatorCorreo(data.Correo, logger, c)
		if valid {
			c.JSON(200, "OK")
		}
	}

}
func EstadoPostulacion(c *gin.Context, db *sql.DB, logger *logrus.Entry) (data models.Registro_Postulacion) {
	var registro models.Registro_Postulacion
	var postulacion models.Postulacion
	cant_aceptados := function.Cantidad_aceptados(db, registro.Rut, registro, logger) //llame a la funcion sql
	if cant_aceptados == 0 {
		function.Postulacion_approved(db, registro.Rut, registro.Electivo, registro, postulacion, logger)
	}
	if cant_aceptados == 1 {
		if registro.Estado == false && registro.Cantidad_Electivos == 2 {
			function.Postulacion_approved(db, registro.Rut, registro.Electivo, registro, postulacion, logger)
		} else {
			function.Postulacion_rejected(db, registro.Rut, registro.Electivo, registro, postulacion, logger)
		}
	}
	if cant_aceptados == 2 {
		function.Postulacion_rejected(db, registro.Rut, registro.Electivo, registro, postulacion, logger)
	}

	return
}
