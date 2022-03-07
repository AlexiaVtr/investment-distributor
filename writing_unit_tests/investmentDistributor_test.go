package main_test

import (
	main "investmentsDistributor"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Function.go
func TestGetCredit(t *testing.T) {
	a, b, c, err := main.GetCredit(3000)

	if err != nil {
		t.Error("TestGetCredit no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestGetCredit ha funcionado correctamente.\n Entregó los valores:", a, b, c)
	}

}

func TestMakeCredit(t *testing.T) {
	a, b, c := main.MakeCredit(3000, 300, 500, 700)

	if a == 0 && b == 0 && c == 0 {
		t.Error("TestMakeCredit no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestMakeCredit ha funcionado correctamente.\n Entregó los valores:", a, b, c)
	}

}

func TestGetAverage(t *testing.T) {
	a := main.GetAverage(10, 2)
	if a != 5 {
		t.Error("TestGetAverage no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestGetAverage ha funcionado correctamente.")
	}
}

func TestCalculateAverage(t *testing.T) {
	err := main.CalculateAverage()

	if err != nil {
		t.Error("TestCalculateAverage no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestCalculateAverage ha funcionado correctamente.")
	}
}

func TestGetStatisticsData(t *testing.T) {
	a, err := main.GetStatisticsData()

	if err != nil {
		t.Error("TestGetStatisticsData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestGetStatisticsData ha funcionado correctamente.\nCon el valor de:", a)
	}
}

func TestSetStatisticsData(t *testing.T) {
	a, err := main.GetStatisticsData()
	err = main.SetStatisticsData(a)

	if err != nil {
		t.Error("TestSetStatisticsData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestSetStatisticsData ha funcionado correctamente.")
	}
}

func TestPutStatisticsData(t *testing.T) {
	a, err := main.GetStatisticsData()
	err = main.PutStatisticsData(a)

	if err != nil {
		t.Error("TestPutStatisticsData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestPutStatisticsData ha funcionado correctamente.")
	}

}

func TestGetInvestmentData(t *testing.T) {
	a, err := main.GetInvestmentData()

	if err != nil {
		t.Error("TestGetInvestmentData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestGetInvestmentData ha funcionado correctamente.\nCon el valor de:", a)
	}

}

func TestSetInvestmentData(t *testing.T) {
	a, err := main.GetInvestmentData()
	err = main.SetInvestmentData(a)

	if err != nil {
		t.Error("TestSetInvestmentData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestSetInvestmentData ha funcionado correctamente.")
	}
}

func TestPutInvestmentData(t *testing.T) {
	a, err := main.GetInvestmentData()
	err = main.PutInvestmentData(a)

	if err != nil {
		t.Error("TestPutInvestmentData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestPutInvestmentData ha funcionado correctamente.")
	}
}

func TestDeleteData(t *testing.T) {
	s, _ := main.GetStatisticsData()
	a, _ := main.GetInvestmentData()
	s, a = main.DeleteData(s, a)
	if s.Average_successful_investment != 0 && s.Average_unsuccessful_investment != 0 &&
		s.Total_assignments_made != 0 && s.Total_successful_assignments != 0 &&
		s.Total_unsuccessful_assignments != 0 && a.Negative != 0 && a.Positive != 0 {
		t.Error("TestDeleteData no ha funcionado.")
		t.Fail()
	} else {
		t.Log("TestDeleteData ha funcionado correctamente.")
	}

}

func TestDBConnection(t *testing.T) {
	_, err := main.DBConnection()

	if err != nil {
		t.Error("TestDBConnection no ha funcionado: Se obtuvo", err)
	} else {
		t.Log("TestDBConnection ha funcionado correctamente.")
	}
}

//Handlers.go

func TestHandleCreditAssignment(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/credit-assignment", nil)
	res := httptest.NewRecorder()

	main.HandleCreditAssignment(res, req)

	// Se espera un BadRequest ya que por consigna si no hay data se debe responder con un 400:
	if res.Code != http.StatusBadRequest {
		t.Errorf("TestHandleCreditAssignment no ha funcionado: Se esperaba %d se obtuvo %d", res.Code, http.StatusOK)
	} else {
		t.Log("TestHandleCreditAssignment ha funcionado correctamente.")
	}
}

func TestHandleStatistics(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/assignment", nil)
	res := httptest.NewRecorder()

	main.HandleStatistics(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("TestHandleStatistics no ha funcionado: Se esperaba %d se obtuvo %d", res.Code, http.StatusOK)
	} else {
		t.Log("TestHandleStatistics ha funcionado correctamente.")
	}
}

func TestHandleDeleteStatistics(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/assignment", nil)
	res := httptest.NewRecorder()

	main.HandleDeleteStatistics(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("TestHandleDeleteStatistics no ha funcionado: Se esperaba %d se obtuvo %d", res.Code, http.StatusOK)
	} else {
		t.Log("TestHandleDeleteStatistics ha funcionado correctamente.")
	}
}
