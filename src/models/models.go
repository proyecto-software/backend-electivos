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
	Id          int    `json:"value"`
	Nombre      string `json:"label"`
	NRC         string `json:"-"`
	Id_profesor int    `json:"-"`
}

type Postulacion struct {
	Id          int  `json:"id_Postulacion"`
	Cantidad    int  `json:"cantidad_ramos"`
	Id_electivo int  `json:"id_Electivo"`
	Aprobado    bool `json:"aprovado"`
}
type Solicitud struct {
	Id               int `json:"id_Solicitud"`
	Id_Postulacion_1 int `json:"id_Postulacion_1"`
	Id_Postulacion_2 int `json:"id_Postulacion_2"`
	Id_Postulacion_3 int `json:"id_Postulacion_3"`
	Id_alumno        int `json:"id_alumno"`
}
type Alumno struct {
	Id                  int    `json:"id_alumno"`
	Rut                 string `json:"rut"`
	Nombre              string `json:"nombre"`
	Correo              string `json:"correo"`
	Id_carrera          int    `json:"id_carrera"`
	Semestre_incompleto int    `json:"semetre_incompleto"`
	Cantidad_ramos      int    `json:"cantidad_ramos"`
}
type Carrera struct {
	Id     int    `json:"value"`
	Nombre string `json:"labels"`
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
