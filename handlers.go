package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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
		PutInvestmentData(average)
		// Se almacena la inversión completa para luego sacar el promedio:
		average.Negative += int64(investmentAmount.Investment)

		//Respuesta:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		// De lo contrario se transforma la resp en JSON y se retorna con 200:
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)

		// Adicional se almacena cantidad e inversión en la variable statistics:
		statisticsData.Total_successful_assignments += 1
		statisticsData.Total_assignments_made += 1
		average.Positive += int64(investmentAmount.Investment)
		log.Print(statisticsData)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	}
	// Se almacenan los datos en Firebase:
	PutInvestmentData(average)
	PutStatisticsData(statisticsData)
	//Se resetean los datos de las variables usadas:
	DeleteData(statisticsData, average)
}

func HandleStatistics(w http.ResponseWriter, r *http.Request) {
	var err error

	// Se calcula y almacena el promedio:
	err = CalculateAverage()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln("HandleStatistics:", err)
	} else {
		// Si no hay errores, retorna la respuesta:
		statisticsData, err = GetStatisticsData()
		data, _ := json.Marshal(statisticsData)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}

func HandleDeleteStatistics(w http.ResponseWriter, r *http.Request) {

	// Borra los datos de las variables de statistics e investment:
	statisticsData, average = DeleteData(statisticsData, average)
	err := SetStatisticsData(statisticsData)
	err = SetInvestmentData(average)
	if err != nil {
		log.Fatalln(err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Print("Se ha borrado la información.")
	}
}
