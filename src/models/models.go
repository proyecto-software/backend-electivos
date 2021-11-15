package models

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
	Id                  int    `json:"Id_alumno"`
	Rut                 string `json:"Rut"`
	Nombre              string `json:"Nombre"`
	Correo              string `json:"Correo"`
	Id_carrera          int    `json:"Id_carrera"`
	Semestre_incompleto int    `json:"Semetre_incompleto"`
	Cantidad_ramos      int    `json:"Cantidad_ramos"`
}
type Carrera struct {
	Id     int    `json:"Id_Carrera"`
	Nombre string `json:"Nombre"`
}

// User defines the user in db
type Admin_user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Alumnoucn struct {

	//Rut del alumno
	Rut string `json:"rut"`
	//Nombre del alumno
	Nombre string `json:"nombre"`
	//Apellido del alumno
	Apellido string `json:"apellido"`
	//Carrera del alumno
	Carrera string
	//El correo institucional del alumno
	Correo string `json:"correo"`
	//El curso con el semestre más retrasado.
	Semestre_incompleto int `json:"semestre_incompleto_menor"`
	//la cantidad de ramos que tiene inscritos en el presente semestre(example: los del primer semestre del 2021)
	Cant_ramos int `json:"cantidad_ramos_inscritos"`
	//La lista con todos los ramos que ha tomado durante su vida universitaria
	Ramos []Ramo `json:"ramos_tomados"` //histórico (rendimiento)

	//Electivo [3]Electivo
}
type Ramo struct {
	//el código del ramo
	Nrc string
	//el nombre del ramo
	Nombre string
	//la nota del correspondiente ramo
	Nota float64
	//las veces que ha dado el alumno ese ramo
	Oportunidad int
	//la cantidad de créditos del ramo
	Creditos int
}
