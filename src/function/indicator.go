package function

import (
	"electivos-ucn/src/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func createIndicador(w http.ResponseWriter, r *http.Request) { //se ingresa un rut y se devuelve su indicador
	vars := mux.Vars(r) //guarda el rut
	var indicador int
	type Alumnos []models.Alumnoucn

	var alumnos = Alumnos{}
	alumnoRut := vars["rut"] //tiene que llamarse igual que en el router

	for _, a := range alumnos {
		if a.Rut == alumnoRut {
			//hace el cálculo del indicador
			//semestre actual (menos capstone) - semestre más atrasado (ej. yo: 9 - 6)
			indicador = 9 - a.Semestre_incompleto
			w.Header().Set("Content-Type", "application/json") //le manda el tipo de dato que va a mandar
			json.NewEncoder(w).Encode(indicador)
		}
	}
}
