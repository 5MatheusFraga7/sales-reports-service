package repository

import (
	"database/sql"
	"fmt"
	"sales-reports-service/internal/db"
	"sales-reports-service/internal/db/adapters"
	"sales-reports-service/internal/models"
)

type SalesRepository struct {
}

func (tr SalesRepository) GetSalesData() []models.Sale {
	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
	}

	rows, err := postgresAdapter.Query(getSalesDataSql())

	if err != nil {
		fmt.Println("Erro ao executar consulta no banco:", err)
	}

	defer rows.Close()
	defer db.CloseConnectionToDatabase(postgresAdapter)

	return mapRowsToSales(rows)
}

func getSalesDataSql() string {
	query := `SELECT * FROM sales limit 100`
	return query
}

func mapRowsToSales(rows *sql.Rows) []models.Sale {
	var sales []models.Sale

	for rows.Next() {
		var sale models.Sale
		err := rows.Scan(&sale.Id, &sale.SelledAt, &sale.Value, &sale.Product)

		if err != nil {
			fmt.Println("Erro ao escanear linha:", err)
		} else {
			sales = append(sales, sale)
		}
	}

	return sales
}
