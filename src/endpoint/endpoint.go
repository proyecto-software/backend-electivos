package endpoint

import (
	"database/sql"
	"electivos-ucn/src/function"
	"electivos-ucn/src/models"
	"strconv"
	"time"

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
		solicitud.Cantidad_Electivos = data.Cantidad
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

		var postulacion models.Postulacion
		//esto tampoco deberia estar
		//
		postulacion.Aprobado = false
		postulacion.Cantidad = data.Cantidad
		if data.Cantidad == 1 {
			postulacion.Id = (time.Now().Second() * time.Now().Nanosecond()) / 10000
			postulacion.Id_electivo = E1.Id
			// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
			solicitud.Id_Postulacion_1 = postulacion.Id
			function.Insert_postulacion(db, postulacion, logger)
		}
		if data.Cantidad == 2 {
			postulacion.Id = (time.Now().Second() * time.Now().Nanosecond()) / 10000
			postulacion.Id_electivo = E1.Id
			// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
			solicitud.Id_Postulacion_1 = postulacion.Id
			function.Insert_postulacion(db, postulacion, logger)
			postulacion.Id = (time.Now().Second() * time.Now().Nanosecond()) / 10000

			postulacion.Id_electivo = E2.Id
			// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
			solicitud.Id_Postulacion_2 = postulacion.Id
			function.Insert_postulacion(db, postulacion, logger)
		}
		if data.Cantidad == 3 {
			postulacion.Id = (time.Now().Second() * time.Now().Nanosecond()) / 10000
			postulacion.Id_electivo = E1.Id
			// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
			solicitud.Id_Postulacion_1 = postulacion.Id
			function.Insert_postulacion(db, postulacion, logger)
			postulacion.Id = (time.Now().Second() * time.Now().Nanosecond()) / 10000
			time.Sleep(time.Second * 1)

			postulacion.Id_electivo = E2.Id
			// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
			solicitud.Id_Postulacion_2 = postulacion.Id
			function.Insert_postulacion(db, postulacion, logger)
			postulacion.Id = (time.Now().Second() * time.Now().Nanosecond()) / 10000
			time.Sleep(time.Second * 1)

			postulacion.Id_electivo = E3.Id
			// AQUI VA UN SELECT POSTULACION RETURN ID.(AUTOINCREMENTO)
			solicitud.Id_Postulacion_3 = postulacion.Id
			function.Insert_postulacion(db, postulacion, logger)

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

		}
		function.Insert_solicitud(db, solicitud, logger)
		logger.Info(Alumno.Nombre)
		c.JSON(200, "OK")
	}

	return
}

func InformeCurricular(c *gin.Context, db *sql.DB, logger *logrus.Entry) {
	rut := c.DefaultQuery("rut", "")
	if function.Validator(rut, logger, c) {
		info := function.Alumno_info(db, rut)
		c.JSON(200, info)
		return
	} else {
		c.JSON(204, nil)
		c.Abort()
		return

	}
}

func TablaInformeCurricular(c *gin.Context, db *sql.DB, logger *logrus.Entry) {
	rut := c.DefaultQuery("rut", "")
	if function.Validator(rut, logger, c) {
		info := function.All_informe_curricular_info(db, rut)
		c.JSON(200, info)
	} else {
		c.JSON(204, nil)
		c.Abort()
		return
	}
	return

}

func RegistroElectivos(c *gin.Context, db *sql.DB, logger *logrus.Entry) (data models.Registro_Electivos) {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	} else {

		Registro := function.Registro_electivos_info(db, data.Año, data.Semestre)
		for i := 0; i < len(Registro); i++ {
			nombre := function.Nombre_electivo(db, Registro[i].Id_electivo)
			logger.Infof(nombre, Registro[i].Cantidad_alumnos)
		}
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
		d = append(d, data_)

	}
	if len(d) != 0 {
		c.JSON(200, d)
	} else {
		ps := models.Registro_Postulacion2{}
		d = append(d, ps)
		c.JSON(204, "{}")
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
