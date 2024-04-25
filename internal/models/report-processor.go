package models

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
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

	if err := writer.WriteAll(records); err != nil {
		panic(err)
	}

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
