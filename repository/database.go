package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Estructura del documento en Firebase:
type Statistics struct {
	Average_successful_investment   float64 `json:"average_successful_investment"`
	Average_unsuccessful_investment float64 `json:"average_unsuccessful_investment"`
	Total_assignments_made          int64   `json:"total_assignments_made"`
	Total_successful_assignments    int64   `json:"total_successful_assignments"`
	Total_unsuccessful_assignments  int64   `json:"total_unsuccessful_assignments"`
}

func main() {

	// Conexión con el SDK:
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()
}
