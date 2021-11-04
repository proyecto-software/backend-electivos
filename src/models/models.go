package models

import (
	"gorm.io/gorm"
)

type Administrador struct {
	Id       int    `json:"Id_administrador"`
	Correo   string `json:"Correo"`
	Rut      string `json:"Rut"`
	Password string `json:"Password"`
}
type Profesor struct {
	Id     int    `json:"Id_profesor"`
	Nombre string `json:"Nombre"`
	Correo string `json:"Correo"`
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
