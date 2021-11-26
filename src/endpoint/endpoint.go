package endpoint

import (
	"database/sql"
	"electivos-ucn/src/function"
	"electivos-ucn/src/models"
	"strconv"

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
		solicitud.Id = 3294
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
			postulacion.Id = i * 3
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
func GetSolicitud(c *gin.Context, db *sql.DB, logger *logrus.Entry) {

	data := function.All_Solicitud_info(db)
	var d []models.Registro_Postulacion2
	var data_ models.Registro_Postulacion2
	for i := 0; i < len(data); i++ {
		if i == 0 {
			data_.Rut = data[i].Rut
			data_.Carrera = data[i].Carrera
			data_.Indicador = data[i].Indicador
			data_.Cantidad_Electivos = data[i].Cantidad_Electivos
			data_.Electivo1 = data[i].Electivo
			data_.Estado1 = data[i].Estado
		} else {
			if data[i].Rut == data[i-0].Rut {
				if data_.Electivo2 == "" {
					data_.Electivo2 = data[i].Electivo
					data_.Estado2 = data[i].Estado
				}
				if data_.Electivo3 == "" {
					data_.Electivo3 = data[i].Electivo
					data_.Estado3 = data[i].Estado
				}
			} else {
				data_.Rut = data[i].Rut
				data_.Carrera = data[i].Carrera
				data_.Indicador = data[i].Indicador
				data_.Cantidad_Electivos = data[i].Cantidad_Electivos
				data_.Electivo1 = data[i].Electivo
				data_.Estado1 = data[i].Estado
			}
		}
	}
	d = append(d, data_)
	c.JSON(200, d)

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
func EstadoPostulacion(c *gin.Context, db *sql.DB, logger *logrus.Entry) (data models.FormatoUpdate) {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
	}

	var postulacion models.Postulacion
	cant_aceptados := function.Cantidad_aceptados(db, data.RutAlumnp, logger) //llame a la funcion sql
	Alumno := function.Alumno_info(db, data.RutAlumnp)
	registro := function.Registro_postulacion_info(db, data.RutAlumnp, data.NombreElectivo)
	Sent := false
	for i := 0; i < len(registro); i++ {
		if !Sent {
			if cant_aceptados == 0 {
				function.Postulacion_approved(db, data.RutAlumnp, data.NombreElectivo, registro[i], postulacion, logger)
				function.SendEmail2(Alumno.Correo, strconv.Itoa((registro[i].Id)), registro[i].Electivo)
				Sent = true

			}
			if cant_aceptados == 1 {
				//ARREGLAR CUANDO SE TIENE 1 ACEPTADO CON 2 ELECTIVOS Y SE QUIERE ELIMINAR ESE ACEPTADO
				if registro[i].Estado == false && registro[i].Cantidad_Electivos == 2 {
					function.Postulacion_approved(db, data.RutAlumnp, data.NombreElectivo, registro[i], postulacion, logger)
					function.SendEmail2(Alumno.Correo, strconv.Itoa((registro[i].Id)), registro[i].Electivo)
					Sent = true
				} else {
					function.Postulacion_rejected(db, data.RutAlumnp, data.NombreElectivo, registro[i], postulacion, logger)
				}
			}
			if cant_aceptados == 2 {
				function.Postulacion_rejected(db, data.RutAlumnp, data.NombreElectivo, registro[i], postulacion, logger)
			}
		}
	}
	return
}
