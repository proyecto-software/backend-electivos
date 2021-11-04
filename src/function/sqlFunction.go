package function

import (
	"database/sql"
	"electivos-ucn/src/models"
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
	rows, err := db.Query("SELECT * FROM public.alumno WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&alumno.Id, &alumno.Rut, &alumno.Nombre, &alumno.Correo, &alumno.Id_carrera)
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

func Electivo_info(db *sql.DB, rut string) (electivo models.Electivo) {
	rows, err := db.Query("SELECT * FROM public.electivo WHERE rut = $1 ", rut)
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
