package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sales-reports-service/internal/db/repository"
	"sales-reports-service/internal/models"
)

type ApiHandler struct {
}

func (p *ApiHandler) HandleInternalServerError(w http.ResponseWriter, r *http.Request) {
	log.Printf("Erro interno do servidor ao processar solicitação para %s %s com parâmetros %v", r.Method, r.URL.Path, r.URL.Query())
	http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
}

func (p *ApiHandler) Get(w http.ResponseWriter, r *http.Request) {

	saleRepository := repository.SalesRepository{}
	sales := saleRepository.GetSalesData()

	processor := models.ReportProcessor{FileName: "Relatório_de_vendas.csv"}
	processor.Call(sales)

	w.Header().Set("Content-Type", "application/json")

	dataResponse := map[string]interface{}{
		"status": "success",
	}
	jsonData, err := json.Marshal(dataResponse)

	if err != nil {
		p.HandleInternalServerError(w, r)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
