package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Realiza la repartición de la inversión a los diferentes créditos:
func GetCredit(inv, credit300, credit500, credit700 int32) (int32, int32, int32, error) {
	n := inv

	// Se realizan todas las restas posibles en 6 casos:
	switch {

	case n > 0:
		credit700, credit500, credit300 = getCredit(n, 700, 500, 300)

		//Si los contadores son ">0" entonces se finalizó la distribución correctamente:
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}

		//Permite que no se saltee los demás casos:
		fallthrough

	case n > 0:

		// El orden de las restas cambia en cada caso:
		credit300, credit500, credit700 = getCredit(n, 300, 500, 700)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit500, credit300, credit700 = getCredit(n, 500, 300, 700)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit500, credit700, credit300 = getCredit(n, 500, 700, 300)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough

	case n > 0:
		credit700, credit300, credit500 = getCredit(n, 700, 300, 500)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}
		fallthrough
	case n > 0:
		credit300, credit700, credit500 = getCredit(n, 300, 700, 500)
		if credit300 != 0 || credit500 != 0 || credit700 != 0 {
			return credit300, credit500, credit700, nil
		}

	}

	// Si ningúno puede dar resto "0" entonces devuelve los contadores y el error creado:
	return credit300, credit500, credit700, &MyError{}
}

// Reutiiza el ciclo for con montos definidos en los parámetros:
func getCredit(i, a, b, c int32) (int32, int32, int32) {

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
func GetAverage(a int64, b int64) float32 {
	var c float32
	c = float32(a) / float32(b)
	return c
}

// Borra los datos de las variables statistics:

func DeleteData() {
	statisticsData.Average_successful_investment = 0
	statisticsData.Average_unsuccessful_investment = 0
	statisticsData.Total_assignments_made = 0
	statisticsData.Total_successful_assignments = 0
	statisticsData.Total_unsuccessful_assignments = 0
	average.Positive = 0
	average.Negative = 0
}

// Almacena la data recibida a specific_statistics en Firebase:
func SetStatisticsData(data Statistics) error {

	// Conexión con el SDK:
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	defer client.Close()

	// Se almacenan los datos modificados en Firebase:
	resultSet, err := client.Collection("statistics").Doc("specific_statistics").Set(context.Background(), data)
	log.Print(resultSet, "almacenado con éxito.")

	return err
}

// Almacena la data recibida a specific_statistics en Firebase:
func SetInvestmentData(data Average) error {

	// Conexión con el SDK:
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	defer client.Close()

	// Se almacenan los datos modificados en Firebase:
	resultSet, err := client.Collection("statistics").Doc("investment").Set(context.Background(), data)
	log.Print(resultSet, "almacenado con éxito.")

	return err
}

// Obtiene los datos de specific_statistics en Firebase:
func GetStatisticsData() (Statistics, error) {
	var dataS Statistics

	// Conexión con el SDK:
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	defer client.Close()

	// Se almacenan los datos modificados en Firebase:
	resultGet, err := client.Collection("statistics").Doc("specific_statistics").Get(context.Background())

	//Se almacena la data en un map:
	response := resultGet.Data()

	// Depuración de la data para almacenar lo obtenido en dataS:
	dataS = MergeData(dataS, response)

	return dataS, err
}

// Obtiene los datos de specific_statistics en Firebase:
func GetInvestmentData() (int64, int64, error) {

	// Conexión con el SDK:
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	defer client.Close()

	// Se almacenan los datos modificados en Firebase:
	resultGet, err := client.Collection("statistics").Doc("investment").Get(context.Background())

	// Se almacena la data en un map:
	response := resultGet.Data()

	// Convertir map en array:
	elementArray := DistributeData(response)

	return elementArray[0], elementArray[1], err
}

// Merge de datos de la BD con statisticsData:
func MergeData(a Statistics, b map[string]interface{}) Statistics {

	// Merge de datos de variables junto con los de BD:
	elementArray := DistributeData(b)

	// Los datos de dataS se suman a los de elementArray.
	// La conversión a float32 es necesaria para coincidir con StatisticsData:
	a.Average_successful_investment = float32(elementArray[0])
	a.Total_assignments_made = elementArray[1]
	a.Total_successful_assignments = elementArray[2]
	a.Average_unsuccessful_investment = float32(elementArray[3])
	a.Total_unsuccessful_assignments = elementArray[4]
	return a
}

// Convierte un map en array:
func DistributeData(a map[string]interface{}) []int64 {
	// Merge de datos de variables junto con los de BD:
	var elementArray []int64

	//Se convierten los elementos string del map a int64 dentro de elementArray:
	for _, elemento := range a {
		newElement, _ := elemento.(int64)
		elementArray = append(elementArray, newElement)
	}

	return elementArray
}
