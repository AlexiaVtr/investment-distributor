package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
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
		statisticsData.Total_unsuccessful_assignments += 1
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		// De lo contrario se transforma la resp en JSON y se retorna 200:
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

		// Además se almacena la cantidad en la variable statistics:
		statisticsData.Total_assignments_made += 1
		statisticsData.Total_successful_assignments += 1
		log.Print(statisticsData)
	}

}

func HandleStatistics(w http.ResponseWriter, r *http.Request) {
	log.Print(statisticsData)
	// Conexión con el SDK:
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	// Envío de la información a la BD:
	result, err := client.Collection("statistics").Doc("specific_statistics").Set(context.Background(), statisticsData)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
	defer client.Close()

}
