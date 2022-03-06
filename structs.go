package main

type Investment int32

type MyError struct {
}

// Request and Responses:
type Amount struct {
	Investment int32 `json:"investment"`
}

type Credits struct {
	Credit_type_300 int32 `json:"credit_type_300"`
	Credit_type_500 int32 `json:"credit_type_500"`
	Credit_type_700 int32 `json:"credit_type_700"`
}

type Statistics struct {
	Average_successful_investment   float32 `json:"average_successful_investment"`
	Average_unsuccessful_investment float32 `json:"average_unsuccessful_investment"`
	Total_assignments_made          int64   `json:"total_assignments_made"`
	Total_successful_assignments    int64   `json:"total_successful_assignments"`
	Total_unsuccessful_assignments  int64   `json:"total_unsuccessful_assignments"`
}
