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
	Id                 int `json:"id_Solicitud"`
	Id_Postulacion_1   int `json:"id_Postulacion_1"`
	Id_Postulacion_2   int `json:"id_Postulacion_2"`
	Id_Postulacion_3   int `json:"id_Postulacion_3"`
	Id_alumno          int `json:"id_alumno"`
	Cantidad_Electivos int `json:"cantidad_electivos"`
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
	Nombre string `json:"label"`
}

// User defines the user in db
type Admin_user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Registro_Postulacion struct {
	Id                 int    `json:"id_registro_postulacion"`
	Rut                string `json:"rut_alumno"`
	Nombre             string `json:"nombre_alumno"`
	Carrera            string `json:"carrera"`
	Indicador          string `json:"indicador"`
	Electivo           string `json:"electivo"`
	Cantidad_Electivos int    `json:"cantidad_electivos"`
	Estado             bool   `json:"estado"`
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
	Id       int    `json:"id_ramo"`
	Nrc      string `json:"nrc"`
	Nombre   string `json:"nombre"`
	Creditos int    `json:"créditos"`
	Semestre int    `json:"semestre"`
}

type Ramo_Alumno struct {
	Id_ramo     int     `json:"id_ramo"`
	Id_alumno   int     `json:"id_alumno"`
	Nota        float64 `json:"nota"`
	Oportunidad int     `json:"oportunidad"`
}

type Registro_Electivos struct {
	Id               int    `json:"id"`
	Nombre           string `json:"nombre"`
	Cantidad_alumnos int    `json:"cantidad_alumnos"`
	Semestre         string `json:"Semestre"` //es string pq el profesor va a mandar con el semestre como ej: Segundo semestre 2020
}

type Informe_Curricular struct {
	Id          int     `json:"id"`
	Rut_alumno  string  `json:"rut_alumno"`
	Nrc         string  `json:"nrc"`
	Nombre_ramo string  `json:"nombre_ramo"`
	Nota        float64 `json:"nota"`
	Oportunidad int     `json:"oportunidad"`
	Semestre    int     `json:"semestre"`
}

type FormatoUpdate struct {
	RutAlumnp      string `json:"rut_alumno"`
	NombreElectivo string `json:"nombre_electivo"`
}

type FormularioAdmin struct {
	Rut_alumno string `json:"rut_alumno"`
}
