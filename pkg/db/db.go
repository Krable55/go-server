package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func Connect() {
	env := os.Getenv("ENV")
	config := GetConfig(env)

	apiKey := config.API_KEY // "Password" from connect menu
	userName := config.USER_NAME
	dbName := config.DATABASE
	port := config.PORT

	fmt.Println("Connecting to database...")

	connectUrl := fmt.Sprintf("postgres://%s:%s@db.bit.io:%v/%s", userName, apiKey, port, dbName)
	conn, err := pgx.Connect(context.Background(), connectUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())

	DB = conn
	os.Setenv("DB", dbName)
	fmt.Printf("Connected to database: %s\n", dbName)
	// Let's see how the census population of Nevada has changed
	// sqlQuery := fmt.Sprintf(`SELECT "Year", "Resident Population" FROM "dliden/2020_Census_Reapportionment"."Historical Apportionment" WHERE "Name" = 'Nevada';`)
	// rows, err := conn.Query(context.Background(), sqlQuery)

	// if err != nil {
	// 	fmt.Println("query error")
	// 	panic(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var Year int
	// 	var Population int

	// 	err = rows.Scan(&Year, &Population)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(Year, Population)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
}
