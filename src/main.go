package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func main() {

	// Open up our database connection.
	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:apgtlmgs1@192.168.1.150:5432/lego")

	// defer the close till after the main function has finished
	// executing
	defer conn.Close(context.Background())
	var greeting string
	//
	conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
}
