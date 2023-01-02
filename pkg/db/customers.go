package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Customers
type Customer struct {
	ID           int    `db:"id" json:"id"`
	Name         string `json:"name"`
	Phone_number int    `json:"phone_number"`
	Address      string `json:"address"`
	Job_id       int    `json:"job_id"`
}

func GetAllCustomers() ([]byte, error) {
	ctx := context.Background()

	var customers []*Customer
	sqlQuery := `SELECT * FROM customers`
	db, err := pgxpool.Connect(ctx, DB.Config().ConnString())
	isConnected := db.Ping(ctx)

	if isConnected != nil {
		fmt.Println("connection error")
		panic(isConnected)
	}

	pgxscan.Select(ctx, db, &customers, sqlQuery)
	if selectErr := pgxscan.Select(ctx, db, &customers, sqlQuery); selectErr != nil {
		fmt.Println("query error")
		panic(selectErr)
	}

	if err != nil {
		fmt.Println("query error")
		panic(err)
	}

	data, jsonErr := json.Marshal(customers)
	if jsonErr != nil {
		fmt.Println("Error converting users to json: ")
		fmt.Println("query error")
		panic(jsonErr)
	}

	return data, nil
}

func GetCustomerById(id int) Customer {
	return Customer{}
}
