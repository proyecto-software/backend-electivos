package function

import (
	"electivos-ucn/src/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func createAlumnos(w http.ResponseWriter, r *http.Request, alumnos []models.Alumnoucn) { //no va a ser necesario en el proyecto
	var newAlumno models.Alumnoucn
	reqBody, err := ioutil.ReadAll(r.Body) //ioutil se usa para el tema del input
	if err != nil {
		fmt.Fprintf(w, "Insert valid data")
	}
	json.Unmarshal(reqBody, &newAlumno)

	alumnos = append(alumnos, newAlumno)

	w.Header().Set("Content-Type", "application/json") //le manda el tipo de dato que va a mandar
	w.WriteHeader(http.StatusCreated)                  //para enviarle que el alumno fue creado
	json.NewEncoder(w).Encode(newAlumno)               //le manda al alumno creado para que vea que se creó
}

func getAlumnos(w http.ResponseWriter, r *http.Request, alumnos []models.Alumnoucn) {
	w.Header().Set("Content-Type", "application/json") //le manda el tipo de dato que va a mandar
	json.NewEncoder(w).Encode(alumnos)

}

func getAlumno(w http.ResponseWriter, r *http.Request, alumnos []models.Alumno) {
	vars := mux.Vars(r) //guarda el rut

	alumnoRut := vars["rut"] //tiene que llamarse igual que en el router

	for _, a := range alumnos {
		if a.Rut == alumnoRut {
			w.Header().Set("Content-Type", "application/json") //le manda el tipo de dato que va a mandar
			json.NewEncoder(w).Encode(a)                       //se le responde con el alumno
		}
	}
}
