package models

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"fmt"
)

type ReportProcessor struct {
	FileName    string
	GeneratedAt time.Time
}

func (rp *ReportProcessor) Call(data []Sale) {
	
	file, err := os.Create(rp.FileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	setHeadersToFile(writer)

	records := make([][]string, len(data))

	for i, row := range data {
		records[i] = []string{row.Product, row.SelledAt.Format("2006-01-02"), strconv.FormatFloat(row.Value, 'f', 2, 64)}
	}

	recordCh := make(chan []string)
	splitedRecords := splitRecords(records)

	// Iniciar a goroutine de escrita
	go func() {
		for record := range recordCh {
				if err := writer.Write(record); err != nil {
						panic(err)
				}
		}
		writer.Flush()
	}()

	// Enviar registros para o canal de forma concorrente
	for _, record := range splitedRecords {
		recordCh <- record
	}

	// Fechar o canal após enviar todos os registros
	close(recordCh)

	if err := writer.Error(); err != nil {
		panic(err)
	}

	writer.Flush()
}

func setHeadersToFile(writer *csv.Writer) {
	headers := []string{"Produto", "data de venda", "Valor da venda"}

	if err := writer.Write(headers); err != nil {
		panic(err)
	}
}

func splitRecords(records [][]string) [][]string {
	totalRecords := len(records)
	numParts := 5
	chunkSize := (totalRecords + numParts - 1) / numParts

	var dividedRecords [][]string

	for i := 0; i < totalRecords; i += chunkSize {
		end := i + chunkSize
		if end > totalRecords {
			end = totalRecords
		}
		
		fmt.Printf("DADO: %v\n", records[i:end])
		dividedRecords = append(dividedRecords, records[i:end]...)
	}

	return dividedRecords
}