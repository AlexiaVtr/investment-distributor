package main_test

import (
	"encoding/json"
	main "investmentsDistributor"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreditAssigment(t *testing.T) {
	//Cliente Cloud Firestore:
	main.ClientF, main.Err = main.DBConnection()
	if main.Err != nil {
		log.Fatal("No se ha podido inicializar la conexión con Firebase:", main.Err)
	}
	defer main.ClientF.Close()

	inv := &main.Amount{
		Investment: 2000,
	}
	dummy, _ := json.Marshal(inv)

	w := httptest.NewRecorder()
	json.NewEncoder(w).Encode(dummy)
	r := httptest.NewRequest("POST", "/credit-assignment", nil)
	main.HandleCreditAssignment(w, r)
	res := w.Result()
	if main.Err != nil && res.StatusCode == http.StatusOK {
		t.Error("TestCreditAssigment no ha funcionado. Se obtuvo:", main.Err)
	} else {
		t.Log("TestCreditAssigment ha funcionado.")
	}

}

func TestStatistics(t *testing.T) {
	//Cliente Cloud Firestore:
	main.ClientF, main.Err = main.DBConnection()
	if main.Err != nil {
		log.Fatal("No se ha podido inicializar la conexión con Firebase:", main.Err)
	}
	defer main.ClientF.Close()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/statistics", nil)
	main.HandleStatistics(w, r)
	res := w.Result()
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Print(bodyString)
	}
	if main.Err != nil && res.StatusCode == http.StatusOK {
		t.Error("TestStatistics no ha funcionado. Se obtuvo:", main.Err)
	} else {
		t.Log("TestStatistics ha funcionado.")
	}

}
