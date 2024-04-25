package main

import (
	"fmt"
	"log"
	"net/http"
	"sales-reports-service/api/routes"
)

func main() {
	router := routes.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "localhost", "3000"),
		Handler: router,
	}

	log.Printf(fmt.Sprintf("Server is running in %s:%s", "localhost", "3000"))
	log.Fatal(server.ListenAndServe())
}
