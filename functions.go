package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//Conexión con el SDK de Firebase:
func DBConnection() (*firestore.Client, error) {
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	return client, err
}

// Realiza la repartición de la inversión a los diferentes créditos:
func GetCredit(inv int32) (int32, int32, int32, error) {
	var credit300, credit500, credit700 int32
	n := inv

	// Se realizan todas las restas posibles en 6 casos:
	switch {

	case n > 0:
		credit700, credit500, credit300 = MakeCredit(n, 700, 500, 300)

		//Si los contadores son ">0" entonces se finalizó la distribución correctamente:
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}

		//Permite que no se saltee los demás casos:
		fallthrough

	case n > 0:

		// El orden de las restas cambia en cada caso:
		credit300, credit500, credit700 = MakeCredit(n, 300, 500, 700)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit500, credit300, credit700 = MakeCredit(n, 500, 300, 700)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit500, credit700, credit300 = MakeCredit(n, 500, 700, 300)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit700, credit300, credit500 = MakeCredit(n, 700, 300, 500)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough
	case n > 0:
		credit300, credit700, credit500 = MakeCredit(n, 300, 700, 500)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}

	}

	// Si ningúno puede dar resto "0" entonces devuelve los contadores y el error creado:
	return credit300, credit500, credit700, &MyError{}
}

// Reutiiza el ciclo for con montos definidos en los parámetros:
func MakeCredit(i, a, b, c int32) (int32, int32, int32) {

	// Las variables son contadores para definir la cantidad de créditos a asignar:
	var creditA, creditB, creditC int32

	// "i"     -> inversión
	// "a,b,c" -> montos
	for i > 0 {
		if i >= a {
			i -= a
			creditA += 1
		}
		if i >= b {
			i -= b
			creditB += 1
		}
		if i >= c {
			i -= c
			creditC += 1
		}

		// Si la inversión recibida no se acredito al 100% finaliza el ciclo:

		if i < 300 && i != 0 {
			creditA, creditB, creditC = 0, 0, 0
			break
		}
	}

	//Retorna los contadores sin importar los resultados:
	return creditA, creditB, creditC

}

// Retorna el promedio de lo enviado:
func GetAverage(a int64, b int64) int64 {
	var c int64
	c = a / b
	log.Print("GetAverage:", a, b, c)
	return c
}

func CalculateAverage() (err error) {
	// Obtener el promedio de inversión.

	// Toma los datos desde la BD:
	statisticsData, err = GetStatisticsData()
	log.Print("CalculateAverage:", statisticsData)
	// Se toma la inversión total almacenada en la bd:
	average, err = GetInvestmentData()

	// Se llama a GetAverage para sacar el promedio:
	if average.Positive > 0 && statisticsData.Total_successful_assignments > 0 {
		statisticsData.Average_successful_investment = GetAverage(average.Positive,
			statisticsData.Total_successful_assignments)
	}
	if average.Negative > 0 && statisticsData.Total_unsuccessful_assignments > 0 {
		statisticsData.Average_unsuccessful_investment = GetAverage(average.Negative,
			statisticsData.Total_unsuccessful_assignments)
	}
	// Envío de la información a la BD con el promedio actualizado:
	err = SetStatisticsData(statisticsData)
	if err != nil {
		log.Fatalln("SetStatisticData:", err)
	}

	// Se elimina la información de las variables:
	statisticsData, average = DeleteData(statisticsData, average)
	log.Print("CalculateAverage:", statisticsData, average)
	return err
}

// Borra los datos de las variables statistics:

func DeleteData(s Statistics, a Average) (Statistics, Average) {
	s.Average_successful_investment = 0
	s.Average_unsuccessful_investment = 0
	s.Total_assignments_made = 0
	s.Total_successful_assignments = 0
	s.Total_unsuccessful_assignments = 0
	a.Positive = 0
	a.Negative = 0
	log.Print("Datos eliminados correctamenre:", s, a)

	return s, a
}

// Almacena la data recibida a specific_statistics en Firebase:
func PutStatisticsData(data Statistics) error {
	//Merge de los datos ingresados y los datos de la BD:
	var databd Statistics
	databd, Err = GetStatisticsData()
	data.Total_assignments_made += databd.Total_assignments_made
	data.Total_successful_assignments += databd.Total_successful_assignments
	data.Total_unsuccessful_assignments += databd.Total_unsuccessful_assignments
	// Se almacenan los datos modificados en Firebase:
	resultSet, err := ClientF.Collection("statistics").Doc("specific_statistics").Set(context.Background(), data)
	log.Print(resultSet, "\nspecific_stadistics almacenadas con éxito.")

	return err
}

// Reemplaza la bd con la data recibida en specific_stadistic en Firebase:
func SetStatisticsData(data Statistics) error {
	// Se almacenan los datos modificados en Firebase:
	resultSet, err := ClientF.Collection("statistics").Doc("specific_statistics").Set(context.Background(), data)
	log.Print(resultSet, "\nspecific_stadistics enviadas con éxito.")

	return err
}

// Almacena la data recibida a specific_statistics en Firebase:
func PutInvestmentData(data Average) error {
	//Merge de los datos ingresados y los datos de la BD:
	var databd Average
	databd, Err = GetInvestmentData()
	data.Negative += databd.Negative
	data.Positive += databd.Positive
	// Se almacenan los datos modificados en Firebase:
	resultSet, err := ClientF.Collection("statistics").Doc("investment").Set(context.Background(), data)
	log.Print(resultSet, data, "investment almacenado con éxito.")

	return err
}

// Reemplaza la bd con la data recibida en investment en Firebase:
func SetInvestmentData(data Average) error {
	// Se almacenan los datos modificados en Firebase:
	resultSet, err := ClientF.Collection("statistics").Doc("investment").Set(context.Background(), data)
	log.Print(resultSet, data, "investment enviado con éxito.")

	return err
}

// Obtiene los datos de specific_statistics en Firebase:
func GetStatisticsData() (Statistics, error) {
	var dataS Statistics
	// Se almacenan los datos de la BD:
	resultGet, err := ClientF.Collection("statistics").Doc("specific_statistics").Get(context.Background())

	//Se almacena la data en un map:
	response := resultGet.Data()

	// Conversión de la data para almacenar lo obtenido en dataS:
	data, _ := json.Marshal(response)
	json.Unmarshal(data, &dataS)

	return dataS, err
}

// Obtiene los datos de specific_statistics en Firebase:
func GetInvestmentData() (Average, error) {
	var dataS Average
	// Se almacenan los datos modificados en Firebase:
	resultGet, err := ClientF.Collection("statistics").Doc("investment").Get(context.Background())

	// Se almacena la data en un map:
	response := resultGet.Data()

	// Conversión de la data para almacenar lo obtenido en dataS:
	data, _ := json.Marshal(response)
	json.Unmarshal(data, &dataS)

	return dataS, err
}
