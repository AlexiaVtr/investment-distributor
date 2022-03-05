package main

import (
	"errors"
	"fmt"
	"os"
)

// Interfaces
type CreditAssing interface {
	Assing(investment int32) (int32, int32, int32, error)

	// Método que imprime resultado:
	PrintCredit(inv, credit300, credit500, credit700 int32)
}

//Manejo de errores
type error interface {
	Error() string
}

//Structs
type Investment int32

type MyError struct {
}

// Request and Responses Structs:
type Amount struct {
	Investment int32 `json:"investment"`
}

type Credits struct {
	Credit_type_300 int32 `json:"credit_type_300"`
	Credit_type_500 int32 `json:"credit_type_500"`
	Credit_type_700 int32 `json:"credit_type_700"`
}

//Methods

// Control de error de distribución:
func (m *MyError) Error() string {
	return "El monto de inversión no se puede distribuir correctamente."
}

// Obtención de los créditos según inversión:
func (i Investment) Assing(n int32) (int32, int32, int32, error) {
	var credit300, credit500, credit700 int32
	if n != 0 {
		return GetCredit(n, credit300, credit500, credit700)

	} else {
		return credit300, credit500, credit700, errors.New("La inversión debe ser mayor a cero.")
	}
}

func (i Investment) PrintCredit(inv, credit300, credit500, credit700 int32) {
	fmt.Printf("%d: %d x $300 + %d x $500 + %d x $700 = $%d", inv, credit300, credit500, credit700, inv)
}

//variables globales:
var investmentAmount Amount
var response Credits

func main() {

	//Servidor

	// Puerto dinámico:
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8000"
	}
	server := NewServer(":" + Port)
	fmt.Println("Server listen in port:", Port)
	//Endpoints:

	server.Handle("/", "GET", HandleRoot)
	server.Handle("/credit-assignment", "POST", HandleCreditAssignment)
	//	server.Handle("/statistics", "POST", statistics)

	server.Listen()
}
