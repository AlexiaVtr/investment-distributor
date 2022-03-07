package main_test

import (
	main "investmentsDistributor"
	"testing"
)

func TestDBConnection(t *testing.T) {
	_, err := main.DBConnection()

	if err != nil {
		t.Error("TestDBConnection no ha funcionado: Se obtuvo", err)
	} else {
		t.Log("TestDBConnection ha funcionado correctamente.")
	}
}
