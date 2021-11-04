package models

import (
	"gorm.io/gorm"
)

type Formulario struct {
	Rut       string `json:"rut"`
	Nombre    string `json:"nombre"`
	Correo    string `json:"correo"`
	Carrera   string `json:"carrera"`
	Cantidad  int    `json:"cantidad"`
	Electivo1 string `json:"electivo1"`
	Electivo2 string `json:"electivo2"`
	Electivo3 string `json:"electivo3"`
}

type Administrador struct {
	Id       int    `json:"ed_administrador"`
	Correo   string `json:"correo"`
	Rut      string `json:"rut"`
	Password string `json:"password"`
}
type Profesor struct {
	Id     int    `json:"id_profesor"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

type Electivo struct {
	Id          int    `json:"Id_Electivo"`
	Nombre      string `json:"Nombre"`
	NRC         string `json:"NRC"`
	Id_profesor int    `json:"Id_profesor"`
}

type Postulacion struct {
	Id          int  `json:"Id_Postulacion"`
	Cantidad    int  `json:"Cantidad_ramos"`
	Id_electivo int  `json:"Id_Electivo"`
	Aprobado    bool `json:"Aprovado"`
}
type Solicitud struct {
	Id               int `json:"Id_Solicitud"`
	Id_Postulacion_1 int `json:"Id_Postulacion_1"`
	Id_Postulacion_2 int `json:"Id_Postulacion_2"`
	Id_Postulacion_3 int `json:"Id_Postulacion_3"`
	Id_alumno        int `json:"Id_alumno"`
}
type Alumno struct {
	Id         int    `json:"Id_alumno"`
	Rut        string `json:"Rut"`
	Nombre     string `json:"Nombre"`
	Correo     string `json:"Correo"`
	Id_carrera int    `json:"Id_carrera"`
}
type Carrera struct {
	Id     int    `json:"Id_Carrera"`
	Nombre string `json:"Nombre"`
}

// User defines the user in db
type Admin_user struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
