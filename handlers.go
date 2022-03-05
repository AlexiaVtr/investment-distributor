package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Declaración de variables para los responses:

var Code404 string = "RESPONSE CODE: 404"
var Code200 string = "RESPONSE CODE: 200"

// Handlers de redirección:

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusExpectationFailed)
	fmt.Fprintf(w, Code404)
}

func HandleCreditAssignment(w http.ResponseWriter, r *http.Request) {

	//Variables:
	var i Investment

	//Obtención de requets:
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&investmentAmount)

	//Distribución de la inversión:

	// Se llama al método Assing para obtener los créditos posibles:
	response.Credit_type_300,
		response.Credit_type_500,
		response.Credit_type_700, err = CreditAssing.Assing(i, investmentAmount.Investment)

	// Si hubo un error se almacena y se retorna un 400:
	if err != nil {
		statistics.Total_unsuccessful_assignments += 1
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		// De lo contrario se transforma la resp en JSON y se retorna 200:
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

		// Además se almacena la cantidad en la variable statistics:
		statistics.Total_assignments_made += 1
		statistics.Total_successful_assignments += 1
	}

}
