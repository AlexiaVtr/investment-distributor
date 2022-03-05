package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	var investment int32

	//Obtención de requets:
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&investmentAmount)
	log.Fatal(err)

	//Distribución de la inversión:

	//
	fmt.Println(investmentAmount.Investment)
	investment = investmentAmount.Investment

	// Se llama al método Assing para obtener los créditos posibles:
	response.Credit_type_300,
		response.Credit_type_500,
		response.Credit_type_700, err = CreditAssing.Assing(i, investment)

	// Si hubo un error se almacena y se retorna un 400:
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	} else {
		// De lo contrario se transforma en JSON y retorna:
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

}
