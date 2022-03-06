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
	var i Investment

	//Obtención de requets:
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&investmentAmount)

	//Distribución de la inversión.

	// Se registra la asignación:
	statisticsData.Total_assignments_made += 1
	// Se llama al método Assing para obtener los créditos posibles:
	response.Credit_type_300,
		response.Credit_type_500,
		response.Credit_type_700, err = CreditAssing.Assing(i, investmentAmount.Investment)

	// Si hubo un error se almacenan los datos para enviar a la BD y se retorna un 400:
	if err != nil {
		statisticsData.Total_unsuccessful_assignments += 1

		// Se almacena la inversión completa para luego sacar el promedio:
		average.Negative += int64(investmentAmount.Investment)

		//Respuesta:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		// De lo contrario se transforma la resp en JSON y se retorna con 200:
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

		// Adicional se almacena cantidad e inversión en la variable statistics:
		statisticsData.Total_successful_assignments += 1
		average.Positive += int64(investmentAmount.Investment)
		log.Print(statisticsData)
	}
	// Se almacenan los datos en Firebase:
	SetInvestmentData(average)
	SetStatisticsData(statisticsData)
	//Se resetean los datos de las variables usadas:
	DeleteData()
}

func HandleStatistics(w http.ResponseWriter, r *http.Request) {
	var err error
	// Obtener los datos desde la BD:
	statisticsData, err = GetStatisticsData()

	if err != nil {
		log.Fatalln(err)
	}

	// Obtener el promedio de inversión:
	// Se toma la inversión total almacenada en la bd:
	average.Negative, average.Positive, err = GetInvestmentData()

	// Se llama a GetAverage para sacar el promedio:
	if average.Positive > 0 {
		statisticsData.Average_successful_investment = GetAverage(average.Positive,
			statisticsData.Total_successful_assignments)
	}
	if average.Negative > 0 {
		statisticsData.Average_unsuccessful_investment = GetAverage(average.Negative,
			statisticsData.Total_unsuccessful_assignments)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(err)
	} else {

		// Envío de la información a la BD con el promedio actualizado:
		err = SetStatisticsData(statisticsData)
		if err != nil {
			log.Fatalln(err)
		}

		// Si no hay errores, retorna la respuesta:
		data, _ := json.Marshal(statisticsData)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

}

func HandleDeleteStatistics(w http.ResponseWriter, r *http.Request) {

	// Borra los datos de las variables de statistics e investment:
	DeleteData()
	err := SetStatisticsData(statisticsData)
	if err != nil {
		log.Fatalln(err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Print("Se ha borrado la información.")
	}
}
