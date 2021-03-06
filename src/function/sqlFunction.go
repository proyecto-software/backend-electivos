package function

import (
	"database/sql"
	"electivos-ucn/src/models"
	"fmt"
	"strconv"

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
	rows, err := db.Query("SELECT  a.id,rut,a.nombre,correo, c.nombre carrera, COALESCE(semestre_incompleto,0) FROM public.alumno a join carrera c on a.id_carrera = c.id where rut = $1 ", rut)
	if err != nil {
		panic(err)
	}
	//defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&alumno.Id, &alumno.Rut, &alumno.Nombre, &alumno.Correo, &alumno.Nombre_carrera, &alumno.Semestre_incompleto)
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
	rows, err := db.Query("SELECT id,rut,nombre,carrera,indicador,estado,electivo FROM public.registro_postulacion where rut=$1 and electivo =$2 ", rut, electivo)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reg_post models.Registro_Postulacion
		err = rows.Scan(&reg_post.Id, &reg_post.Rut, &reg_post.Nombre, &reg_post.Carrera, &reg_post.Indicador, &reg_post.Estado, &reg_post.Electivo)
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
	informes = []models.Informe_Curricular{}
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
		err = rows.Scan(&reg_elec.Id, &reg_elec.Nombre, &reg_elec.Cantidad_alumnos, &reg_elec.A??o, &reg_elec.Semestre)
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

func Registro_electivos_info(db *sql.DB, a??o int, semestre int) (reg_elecs []models.Registro_Electivos) {
	rows, err := db.Query("SELECT re.id, e.nombre, re.cantidad_alumnos, re.a??o, re.semestre from (SELECT * FROM public.registro_electivos WHERE a??o = $1 AND semestre = $2) as re inner join public.electivo e on e.id = re.id_electivo", a??o, semestre)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reg_elec models.Registro_Electivos
		err = rows.Scan(&reg_elec.Id, &reg_elec.Nombre, &reg_elec.Cantidad_alumnos, &reg_elec.A??o, &reg_elec.Semestre)
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
func All_Solicitud_info(db *sql.DB) (solicitudes []models.Registro_Postulacion2) {
	rows, err := db.Query("select s.id, a.rut, c.nombre carrera, s.cantidad_electivos ,postulacion1.electivo, postulacion1.estado,postulacion2.electivo, postulacion2.estado,postulacion3.electivo, postulacion3.estado from solicitud s inner join (select p.id, e.nombre electivo, p.aprobado estado from postulacion p inner join electivo e on p.id_electivo = e.id ) postulacion1 on postulacion1.id = s.id_postulacion_1 inner join (select p.id, e.nombre electivo, p.aprobado estado from postulacion p inner join electivo e on p.id_electivo = e.id ) postulacion2 on postulacion2.id = s.id_postulacion_2 inner join (select p.id, e.nombre electivo, p.aprobado estado from postulacion p inner join electivo e on p.id_electivo = e.id ) postulacion3 on postulacion3.id = s.id_postulacion_3 inner join alumno a on a.id = s.id_alumno inner join carrera c on c.id = a.id_carrera")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		var solicitud models.Registro_Postulacion2
		err = rows.Scan(&solicitud.Id, &solicitud.Rut, &solicitud.Carrera, &solicitud.Cantidad_Electivos, &solicitud.Electivo1, &solicitud.Estado1, &solicitud.Electivo2, &solicitud.Estado2, &solicitud.Electivo3, &solicitud.Estado3)
		if err != nil {
			panic(err)
		} else {
			Alumno := Alumno_info(db, solicitud.Rut)
			solicitud.Indicador = strconv.Itoa(9 - Alumno.Semestre_incompleto)
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

func PostulacionUpdate(db *sql.DB, estado bool, rut string, electivo string, logger *logrus.Entry) {

	approved2, e2 := db.Query(`UPDATE postulacion SET aprobado = $1 WHERE (id = (SELECT id_postulacion_1 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $2)) OR id = (SELECT id_postulacion_2 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $2)) OR id = (SELECT id_postulacion_3 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $2)) ) AND id_electivo = (SELECT id FROM electivo WHERE nombre = $3)`, estado, rut, electivo)
	if e2 != nil {
		logger.Infof("Error cambiando el estado")
		recoverError()

	} else {
		logger.Infof("estado de la postulaci??n Cargado con Exito")
	}
	defer approved2.Close()
}

func Postulacion_rejected(db *sql.DB, rut string, electivo string, registro_postulacion models.Registro_Postulacion, postulacion models.Postulacion, logger *logrus.Entry) {
	rejected1, e := db.Query(`UPDATE registro_postulacion SET estado = false WHERE rut = $1 and electivo = $2 `, rut, electivo)
	rejected2, e2 := db.Query(`UPDATE postulacion SET aprobado = false WHERE (id = (SELECT id_postulacion_1 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) OR id = (SELECT id_postulacion_2 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) OR id = (SELECT id_postulacion_3 FROM solicitud WHERE id_alumno = (SELECT id FROM alumno WHERE rut = $1)) ) AND id_electivo = (SELECT id FROM electivo WHERE nombre = $2)`, rut, electivo)
	if e != nil || e2 != nil {
		logger.Infof("Error cambiando el estado")
		recoverError()
	} else {
		logger.Infof("estado de la postulaci??n Cargado con Exito")
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
	//defer db.Close()
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
	//defer db.Close()
}
func recoverError() {
	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			fmt.Printf("")
		}
	}()
	panic("Error")
}

func Nombre_electivo(db *sql.DB, id_electivo int) string {
	var nombre string
	rows, err := db.Query("SELECT nombre FROM public.electivo WHERE id = $1 ", id_electivo)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&nombre)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return nombre

}

func AnoSemestres(db *sql.DB) (informes []models.AnoSemestre) {
	informes = []models.AnoSemestre{}
	rows, err := db.Query("SELECT distinct a??o, semestre FROM public.registro_electivos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var informe models.AnoSemestre
		err = rows.Scan(&informe.Ano, &informe.Semestre)
		if err != nil {
			panic(err)
		} else {
			str_final := fmt.Sprintf("%d-%d", informe.Ano, informe.Semestre)
			informe.Info = str_final
			informe.Info2 = str_final
			informes = append(informes, informe)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return

}
