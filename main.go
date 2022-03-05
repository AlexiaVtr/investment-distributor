package main

import (
	"errors"
	"fmt"
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

func main() {
	var i Investment
	var investment int32

	fmt.Println("Inserte la inversión:")
	fmt.Scanln(&investment)

	// Se llama al método Assing para obtener los créditos posibles:
	credit300, credit500, credit700, err := CreditAssing.Assing(i, investment)

	// Si hubo un error se imprime:
	if err != nil {
		fmt.Println(err)
	} else {
		// De lo contrario se responde con el método PrintCredit:
		CreditAssing.PrintCredit(i, investment, credit300, credit500, credit700)

	}
}
