package function

import (
	"database/sql"
	"electivos-ucn/src/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

func Administrador_info(db *sql.DB, rut string) (admin models.Administrador) {
	rows, err := db.Query("SELECT * FROM public.administrador WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&admin.Id, &admin.Correo, &admin.Rut, &admin.Password)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Alumno_info(db *sql.DB, rut string) (alumno models.Alumno) {
	rows, err := db.Query("SELECT id,rut,nombre,correo,id_carrera,COALESCE(semestre_incompleto,0) ,COALESCE(cantidad_ramos,0) FROM public.alumno WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&alumno.Id, &alumno.Rut, &alumno.Nombre, &alumno.Correo, &alumno.Id_carrera, &alumno.Semestre_incompleto, &alumno.Cantidad_ramos)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Carrera_info(db *sql.DB, rut string) (carrera models.Carrera) {
	rows, err := db.Query("SELECT * FROM public.carrera WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&carrera.Id, &carrera.Nombre)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}
func All_carrera_info(db *sql.DB) (carreras []models.Carrera) {

	rows, err := db.Query("SELECT * FROM public.carrera ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var carrera models.Carrera
		err = rows.Scan(&carrera.Id, &carrera.Nombre)
		if err != nil {
			panic(err)
		} else {
			carreras = append(carreras, carrera)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Electivo_info(db *sql.DB, nombre string) (electivo models.Electivo) {
	rows, err := db.Query("SELECT * FROM public.electivo WHERE nombre = $1 ", nombre)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&electivo.Id, &electivo.Nombre, &electivo.NRC, &electivo.Id_profesor)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}
func All_electivo_info(db *sql.DB) (electivos []models.Electivo) {
	rows, err := db.Query("SELECT * FROM public.electivo")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var electivo models.Electivo
		err = rows.Scan(&electivo.Id, &electivo.Nombre, &electivo.NRC, &electivo.Id_profesor)
		if err != nil {
			panic(err)
		} else {
			electivos = append(electivos, electivo)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func All_registro_postulacion_info(db *sql.DB) (reg_posts []models.Registro_Postulacion) {
	rows, err := db.Query("SELECT * FROM public.registro_postulacion")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reg_post models.Registro_Postulacion
		err = rows.Scan(&reg_post.Id, &reg_post.Rut, &reg_post.Nombre, &reg_post.Carrera, &reg_post.Indicador, &reg_post.Electivo)
		if err != nil {
			panic(err)
		} else {
			reg_posts = append(reg_posts, reg_post)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Postulacion_info(db *sql.DB, rut string) (postulacion models.Postulacion) {
	rows, err := db.Query("SELECT * FROM public.postulacion WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&postulacion.Id, &postulacion.Cantidad, &postulacion.Id_electivo, &postulacion.Aprobado)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Profesor_info(db *sql.DB, rut string) (profesor models.Profesor) {
	rows, err := db.Query("SELECT * FROM public.alumno WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&profesor.Id, &profesor.Nombre, &profesor.Correo)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Solicitud_info(db *sql.DB, rut string) (solicitud models.Solicitud) {
	rows, err := db.Query("SELECT * FROM public.alumno WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&solicitud.Id, &solicitud.Id_alumno, &solicitud.Id_Postulacion_1, &solicitud.Id_Postulacion_2, &solicitud.Id_Postulacion_3)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func Insert_postulacion(db *sql.DB, postulacion models.Postulacion, logger *logrus.Entry) {
	insertDynStmt := `INSERT INTO public.postulacion
		VALUES ($1, $2, $3, $4);`
	_, e := db.Exec(insertDynStmt, &postulacion.Id, &postulacion.Cantidad, &postulacion.Id_electivo, &postulacion.Aprobado)
	if e != nil {
		logger.Infof("Error postulacion")
		recoverError()
	} else {
		logger.Infof("postulacion Cargada con Exito")
	}
}
func Insert_solicitud(db *sql.DB, postulacion models.Solicitud, logger *logrus.Entry) {
	insertDynStmt := `INSERT INTO public.solicitud
		VALUES ($1, $2, $3, $4, $5);`
	_, e := db.Exec(insertDynStmt, &postulacion.Id, &postulacion.Id_alumno, &postulacion.Id_Postulacion_1, &postulacion.Id_Postulacion_2, &postulacion.Id_Postulacion_3)
	if e != nil {
		logger.Infof("Error solicitud")
		recoverError()
	} else {
		logger.Infof("solicitud Cargada con Exito")
	}
}
func recoverError() {
	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			fmt.Printf("")
		}
	}()
	panic("Error")
}
