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
	rows, err := db.Query("SELECT id,rut,nombre,correo,id_carrera,COALESCE(semestre_incompleto,0) FROM public.alumno WHERE rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&alumno.Id, &alumno.Rut, &alumno.Nombre, &alumno.Correo, &alumno.Id_carrera, &alumno.Semestre_incompleto)
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
		err = rows.Scan(&reg_post.Id, &reg_post.Rut, &reg_post.Nombre, &reg_post.Carrera, &reg_post.Indicador, &reg_post.Electivo, &reg_post.Cantidad_Electivos, &reg_post.Estado)
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
func Registro_postulacion_info(db *sql.DB, rut, electivo string) (reg_posts []models.Registro_Postulacion) {
	rows, err := db.Query("SELECT id,rut,nombre,carrera,indicador,estado FROM public.registro_postulacion where rut=$1 and electivo =$2 ", rut, electivo)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reg_post models.Registro_Postulacion
		err = rows.Scan(&reg_post.Id, &reg_post.Rut, &reg_post.Nombre, &reg_post.Carrera, &reg_post.Indicador, &reg_post.Estado)
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

func All_informe_curricular_info(db *sql.DB, rut string) (informes []models.Informe_Curricular) {
	rows, err := db.Query("SELECT * FROM public.informe_curricular WHERE rut = $1", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var informe models.Informe_Curricular
		err = rows.Scan(&informe.Id, &informe.Rut, &informe.Nrc, &informe.Nombre_ramo, &informe.Nota, &informe.Oportunidad, &informe.Semestre)
		if err != nil {
			panic(err)
		} else {
			informes = append(informes, informe)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}

func All_registro_electivos_info(db *sql.DB) (reg_elecs []models.Registro_Electivos) {
	rows, err := db.Query("SELECT * FROM public.registro_electivos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reg_elec models.Registro_Electivos
		err = rows.Scan(&reg_elec.Id, &reg_elec.Nombre, &reg_elec.Cantidad_alumnos, &reg_elec.Año, &reg_elec.Semestre)
		if err != nil {
			panic(err)
		} else {
			reg_elecs = append(reg_elecs, reg_elec)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}

func Registro_electivos_info(db *sql.DB, año int, semestre int) (reg_elecs []models.Registro_Electivos) {
	rows, err := db.Query("SELECT * FROM public.registro_electivos WHERE año = $1 AND semestre = $2", año, semestre)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reg_elec models.Registro_Electivos
		err = rows.Scan(&reg_elec.Id, &reg_elec.Nombre, &reg_elec.Cantidad_alumnos, &reg_elec.Año, &reg_elec.Semestre)
		if err != nil {
			panic(err)
		} else {
			reg_elecs = append(reg_elecs, reg_elec)
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

func Solicitud_info(db *sql.DB, id string) (solicitud models.Registro_Postulacion) {
	rows, err := db.Query("SELECT a.rut,r.carrera,r.indicador,r.electivo,r.estado FROM public.Registro_Postulacion as r inner join alumno a on a.rut =r.rut WHERE id_alumno = $1 ", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&solicitud.Rut, &solicitud.Carrera, &solicitud.Indicador, &solicitud.Electivo, &solicitud.Estado)
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
func All_Solicitud_info(db *sql.DB) (solicitudes []models.Registro_Postulacion) {
	rows, err := db.Query("SELECT a.rut,r.carrera,r.indicador,r.electivo,r.estado FROM public.Registro_Postulacion as r inner join alumno a on a.rut =r.rut ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		var solicitud models.Registro_Postulacion
		err = rows.Scan(&solicitud.Rut, &solicitud.Carrera, &solicitud.Indicador, &solicitud.Electivo, &solicitud.Estado)
		if err != nil {
			panic(err)
		} else {
			solicitud.IDGen = i
			solicitudes = append(solicitudes, solicitud)
			i = i + 1
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}

func Postulacion_approved(db *sql.DB, rut string, electivo string, registro_postulacion models.Registro_Postulacion, postulacion models.Postulacion, logger *logrus.Entry) {
	approved1, e := db.Query(`UPDATE registro_postulacion SET estado = true WHERE rut = $1 and electivo = $2 `, rut, electivo)
	approved2, e2 := db.Query(`UPDATE postulacion SET aprobado = true WHERE (id = (SELECT id_postulacion_1 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) OR id = (SELECT id_postulacion_2 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) OR id = (SELECT id_postulacion_3 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) ) AND id_electivo = (SELECT id FROM electivo WHERE nombre = $2)`, rut, electivo)
	if e != nil || e2 != nil {
		logger.Infof("Error cambiando el estado")
		recoverError()
	} else {
		logger.Infof("estado de la postulación Cargado con Exito")
	}
	defer approved1.Close()
	defer approved2.Close()
}

func Postulacion_rejected(db *sql.DB, rut string, electivo string, registro_postulacion models.Registro_Postulacion, postulacion models.Postulacion, logger *logrus.Entry) {
	rejected1, e := db.Query(`UPDATE registro_postulacion SET estado = false WHERE rut = $1 and electivo = $2 `, rut, electivo)
	rejected2, e2 := db.Query(`UPDATE postulacion SET aprobado = false WHERE (id = (SELECT id_postulacion_1 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) OR id = (SELECT id_postulacion_2 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) OR id = (SELECT id_postulacion_3 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) ) AND id_electivo = (SELECT id FROM electivo WHERE nombre = $2)`, rut, electivo)
	if e != nil || e2 != nil {
		logger.Infof("Error cambiando el estado")
		recoverError()
	} else {
		logger.Infof("estado de la postulación Cargado con Exito")
	}

	defer rejected1.Close()
	defer rejected2.Close()
}

func Cantidad_aceptados(db *sql.DB, rut string, logger *logrus.Entry) int {
	var cantidad int
	rows, err := db.Query("SELECT COUNT(*) FROM public.registro_postulacion WHERE public.registro_postulacion.estado = true AND rut = $1", rut)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&cantidad)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	cantidad += 1
	return cantidad
}

func MaximoID_postulacion(db *sql.DB, logger *logrus.Entry) int {
	var cantidad int
	rows, err := db.Query("SELECT MAX(public.postulacion.id) FROM public.postulacion")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&cantidad)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	cantidad += 1
	return cantidad
}
func Insert_postulacion(db *sql.DB, postulacion models.Postulacion, logger *logrus.Entry) {
	insertDynStmt := `INSERT INTO public.postulacion
		VALUES ($1, $2, $3);`
	_, e := db.Exec(insertDynStmt, &postulacion.Id, &postulacion.Id_electivo, &postulacion.Aprobado)
	if e != nil {
		logger.Infof("Error postulacion")
		recoverError()
	} else {
		logger.Infof("postulacion Cargada con Exito")
	}
}
func Insert_solicitud(db *sql.DB, postulacion models.Solicitud, logger *logrus.Entry, cantidad int) {
	insertDynStmt := `INSERT  INTO public.solicitud (id_alumno,id_postulacion_1,id_postulacion_2,id_postulacion_3,cantidad_electivos)
		VALUES ($1, $2, $3, $4, $5);`
	_, e := db.Exec(insertDynStmt, &postulacion.Id_alumno, &postulacion.Id_Postulacion_1, &postulacion.Id_Postulacion_2, &postulacion.Id_Postulacion_3, &cantidad)
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
