package main

import (
	"sales-reports-service/internal/db"
	"sales-reports-service/internal/db/adapters"
)

func main() {

	// Conexão com banco

	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	if err != nil {
		panic(err)
	}
}
