package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

// Interfaces
type CreditAssing interface {
	Assing(investment int32) (int32, int32, int32, error)
}

//Manejo de errores
type error interface {
	Error() string
}

//Metodos.

// Control de error de distribución:
func (m *MyError) Error() string {
	return "El monto de inversión no se puede distribuir correctamente."
}

// Obtención de los créditos según inversión:
func (i Investment) Assing(n int32) (int32, int32, int32, error) {
	var credit300, credit500, credit700 int32
	if n != 0 {
		return GetCredit(n)

	} else {
		return credit300, credit500, credit700, errors.New("La inversión debe ser mayor a cero.")
	}
}

func (i Investment) PrintCredit(inv, credit300, credit500, credit700 int32) {
	fmt.Printf("%d: %d x $300 + %d x $500 + %d x $700 = $%d", inv, credit300, credit500, credit700, inv)
}

func main() {

	//Cliente Cloud Firestore:
	client, err = DBConnection()
	if err != nil {
		log.Fatal("No se ha podido inicializar la conexión con Firebase:", err)
	}
	defer client.Close()

	//Servidor

	// Puerto dinámico:
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8000"
	}
	server := NewServer(":" + Port)
	fmt.Println("Server listen in port:", Port)

	//Endpoints:
	server.Handle("/credit-assignment", "POST", HandleCreditAssignment)
	server.Handle("/statistics", "POST", HandleStatistics)
	server.Handle("/statistics", "DELETE", HandleDeleteStatistics)

	server.Listen()

}

//variables globales:
var investmentAmount Amount
var response Credits
var statisticsData Statistics
var average Average
var client *firestore.Client
var err error
